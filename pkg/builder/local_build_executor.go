package builder

import (
	"context"
	"os"
	"path"
	"time"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	re_filesystem "github.com/buildbarn/bb-remote-execution/pkg/filesystem"
	"github.com/buildbarn/bb-remote-execution/pkg/proto/remoteworker"
	runner_pb "github.com/buildbarn/bb-remote-execution/pkg/proto/runner"
	"github.com/buildbarn/bb-remote-execution/pkg/runner"
	"github.com/buildbarn/bb-storage/pkg/cas"
	"github.com/buildbarn/bb-storage/pkg/clock"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type localBuildExecutor struct {
	contentAddressableStorage cas.ContentAddressableStorage
	buildDirectoryCreator     BuildDirectoryCreator
	runner                    runner.Runner
	clock                     clock.Clock
	defaultExecutionTimeout   time.Duration
	maximumExecutionTimeout   time.Duration
	inputRootCharacterDevices map[string]int
}

// NewLocalBuildExecutor returns a BuildExecutor that executes build
// steps on the local system.
func NewLocalBuildExecutor(contentAddressableStorage cas.ContentAddressableStorage, buildDirectoryCreator BuildDirectoryCreator, runner runner.Runner, clock clock.Clock, defaultExecutionTimeout time.Duration, maximumExecutionTimeout time.Duration, inputRootCharacterDevices map[string]int) BuildExecutor {
	return &localBuildExecutor{
		contentAddressableStorage: contentAddressableStorage,
		buildDirectoryCreator:     buildDirectoryCreator,
		runner:                    runner,
		clock:                     clock,
		defaultExecutionTimeout:   defaultExecutionTimeout,
		maximumExecutionTimeout:   maximumExecutionTimeout,
		inputRootCharacterDevices: inputRootCharacterDevices,
	}
}

func (be *localBuildExecutor) createCharacterDevices(inputRootDirectory BuildDirectory) error {
	if err := inputRootDirectory.Mkdir("dev", 0777); err != nil && !os.IsExist(err) {
		return util.StatusWrap(err, "Unable to create /dev directory in input root")
	}
	devDir, err := inputRootDirectory.EnterDirectory("dev")
	defer devDir.Close()
	if err != nil {
		return util.StatusWrap(err, "Unable to enter /dev directory in input root")
	}
	for name, number := range be.inputRootCharacterDevices {
		if err := devDir.Mknod(name, os.ModeDevice|os.ModeCharDevice|0666, number); err != nil {
			return util.StatusWrapf(err, "Failed to create character device %#v", name)
		}
	}
	return nil
}

func (be *localBuildExecutor) Execute(ctx context.Context, filePool re_filesystem.FilePool, instanceName string, request *remoteworker.DesiredState_Executing, executionStateUpdates chan<- *remoteworker.CurrentState_Executing) *remoteexecution.ExecuteResponse {
	response := &remoteexecution.ExecuteResponse{
		Result: &remoteexecution.ActionResult{
			ExecutionMetadata: &remoteexecution.ExecutedActionMetadata{},
		},
	}

	// Timeout handling.
	action := request.Action
	if action == nil {
		attachErrorToExecuteResponse(response, status.Error(codes.InvalidArgument, "Request does not contain an action"))
		return response
	}
	var executionTimeout time.Duration
	if action.Timeout == nil {
		executionTimeout = be.defaultExecutionTimeout
	} else {
		var err error
		executionTimeout, err = ptypes.Duration(action.Timeout)
		if err != nil {
			attachErrorToExecuteResponse(
				response,
				util.StatusWrapWithCode(err, codes.InvalidArgument, "Invalid execution timeout"))
			return response
		}
		if executionTimeout > be.maximumExecutionTimeout {
			attachErrorToExecuteResponse(
				response,
				status.Errorf(
					codes.InvalidArgument,
					"Execution timeout of %s exceeds maximum permitted value of %s",
					executionTimeout,
					be.maximumExecutionTimeout))
			return response
		}
	}

	// Obtain build directory.
	actionDigest, err := digest.NewDigestFromPartialDigest(instanceName, request.ActionDigest)
	if err != nil {
		attachErrorToExecuteResponse(response, util.StatusWrap(err, "Failed to extract digest for action"))
		return response
	}
	command := request.Command
	if command == nil {
		attachErrorToExecuteResponse(response, status.Error(codes.InvalidArgument, "Request does not contain a command"))
		return response
	}
	buildDirectory, buildDirectoryPath, err := be.buildDirectoryCreator.GetBuildDirectory(actionDigest, action.DoNotCache)
	if err != nil {
		attachErrorToExecuteResponse(
			response,
			util.StatusWrap(err, "Failed to acquire build environment"))
		return response
	}
	defer buildDirectory.Close()

	// Install hooks on build directory to capture file creation
	// events.
	buildDirectory.InstallHooks(filePool)

	executionStateUpdates <- &remoteworker.CurrentState_Executing{
		ActionDigest: request.ActionDigest,
		ExecutionState: &remoteworker.CurrentState_Executing_FetchingInputs{
			FetchingInputs: &empty.Empty{},
		},
	}

	// Create input root directory inside of build directory.
	if err := buildDirectory.Mkdir("root", 0777); err != nil {
		attachErrorToExecuteResponse(
			response,
			util.StatusWrap(err, "Failed to create input root directory"))
		return response
	}
	inputRootDirectory, err := buildDirectory.EnterBuildDirectory("root")
	if err != nil {
		attachErrorToExecuteResponse(
			response,
			util.StatusWrap(err, "Failed to enter input root directory"))
		return response
	}
	defer inputRootDirectory.Close()

	inputRootDigest, err := actionDigest.NewDerivedDigest(action.InputRootDigest)
	if err != nil {
		attachErrorToExecuteResponse(
			response,
			util.StatusWrap(err, "Failed to extract digest for input root"))
		return response
	}
	if err := inputRootDirectory.MergeDirectoryContents(ctx, inputRootDigest); err != nil {
		attachErrorToExecuteResponse(response, err)
		return response
	}

	if len(be.inputRootCharacterDevices) > 0 {
		if err := be.createCharacterDevices(inputRootDirectory); err != nil {
			attachErrorToExecuteResponse(response, err)
			return response
		}
	}

	// Create parent directories of output files and directories.
	// These are not declared in the input root explicitly.
	outputHierarchy := NewOutputHierarchy(command.WorkingDirectory)
	for _, outputDirectory := range command.OutputDirectories {
		outputHierarchy.AddDirectory(outputDirectory)
	}
	for _, outputFile := range command.OutputFiles {
		outputHierarchy.AddFile(outputFile)
	}
	for _, outputPath := range command.OutputPaths {
		outputHierarchy.AddPath(outputPath)
	}
	if err := outputHierarchy.CreateParentDirectories(inputRootDirectory); err != nil {
		attachErrorToExecuteResponse(response, err)
		return response
	}

	// Create a directory inside the build directory that build
	// actions may use to store temporary files. This ensures that
	// temporary files are automatically removed when the build
	// action completes. When using FUSE, it also causes quotas to
	// be applied to them.
	if err := buildDirectory.Mkdir("tmp", 0777); err != nil {
		attachErrorToExecuteResponse(
			response,
			util.StatusWrap(err, "Failed to create temporary directory inside build directory"))
		return response
	}

	executionStateUpdates <- &remoteworker.CurrentState_Executing{
		ActionDigest: request.ActionDigest,
		ExecutionState: &remoteworker.CurrentState_Executing_Running{
			Running: &empty.Empty{},
		},
	}

	// Invoke the command.
	ctxWithTimeout, cancelTimeout := be.clock.NewContextWithTimeout(ctx, executionTimeout)
	defer cancelTimeout()
	environmentVariables := map[string]string{}
	for _, environmentVariable := range command.EnvironmentVariables {
		environmentVariables[environmentVariable.Name] = environmentVariable.Value
	}
	runResponse, runErr := be.runner.Run(ctxWithTimeout, &runner_pb.RunRequest{
		Arguments:            command.Arguments,
		EnvironmentVariables: environmentVariables,
		WorkingDirectory:     command.WorkingDirectory,
		StdoutPath:           path.Join(buildDirectoryPath, "stdout"),
		StderrPath:           path.Join(buildDirectoryPath, "stderr"),
		InputRootDirectory:   path.Join(buildDirectoryPath, "root"),
		TemporaryDirectory:   path.Join(buildDirectoryPath, "tmp"),
	})

	// Attach the exit code or execution error.
	if runErr == nil {
		response.Result.ExitCode = runResponse.ExitCode
		response.Result.ExecutionMetadata.AuxiliaryMetadata = runResponse.ResourceUsage
	} else {
		attachErrorToExecuteResponse(response, util.StatusWrap(runErr, "Failed to run command"))
	}

	executionStateUpdates <- &remoteworker.CurrentState_Executing{
		ActionDigest: request.ActionDigest,
		ExecutionState: &remoteworker.CurrentState_Executing_UploadingOutputs{
			UploadingOutputs: &empty.Empty{},
		},
	}

	// Upload command output. In the common case, the stdout and
	// stderr files are empty. If that's the case, don't bother
	// setting the digest to keep the ActionResult small.
	if stdoutDigest, err := be.contentAddressableStorage.PutFile(ctx, buildDirectory, "stdout", actionDigest); err != nil {
		attachErrorToExecuteResponse(response, util.StatusWrap(err, "Failed to store stdout"))
	} else if stdoutDigest.GetSizeBytes() > 0 {
		response.Result.StdoutDigest = stdoutDigest.GetPartialDigest()
	}
	if stderrDigest, err := be.contentAddressableStorage.PutFile(ctx, buildDirectory, "stderr", actionDigest); err != nil {
		attachErrorToExecuteResponse(response, util.StatusWrap(err, "Failed to store stderr"))
	} else if stderrDigest.GetSizeBytes() > 0 {
		response.Result.StderrDigest = stderrDigest.GetPartialDigest()
	}
	if err := outputHierarchy.UploadOutputs(ctx, inputRootDirectory, be.contentAddressableStorage, actionDigest, response.Result); err != nil {
		attachErrorToExecuteResponse(response, err)
	}

	return response
}
