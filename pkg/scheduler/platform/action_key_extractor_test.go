package platform_test

import (
	"context"
	"testing"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/buildbarn/bb-remote-execution/pkg/proto/buildqueuestate"
	"github.com/buildbarn/bb-remote-execution/pkg/scheduler/platform"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/testutil"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestActionKeyExtractor(t *testing.T) {
	keyExtractor := platform.ActionKeyExtractor
	ctx := context.Background()
	instanceName := digest.MustNewInstanceName("hello")

	t.Run("InvalidProperties", func(t *testing.T) {
		_, err := keyExtractor.ExtractKey(ctx, instanceName, &remoteexecution.Action{
			Platform: &remoteexecution.Platform{
				Properties: []*remoteexecution.Platform_Property{
					{Name: "os", Value: "linux"},
					{Name: "arch", Value: "aarch64"},
				},
			},
		})
		testutil.RequireEqualStatus(t, status.Error(codes.InvalidArgument, "Platform properties are not sorted"), err)
	})

	t.Run("Success", func(t *testing.T) {
		key, err := keyExtractor.ExtractKey(ctx, instanceName, &remoteexecution.Action{
			Platform: &remoteexecution.Platform{
				Properties: []*remoteexecution.Platform_Property{
					{Name: "arch", Value: "aarch64"},
					{Name: "os", Value: "linux"},
				},
			},
		})
		require.NoError(t, err)
		testutil.RequireEqualProto(t, &buildqueuestate.PlatformQueueName{
			InstanceNamePrefix: "hello",
			Platform: &remoteexecution.Platform{
				Properties: []*remoteexecution.Platform_Property{
					{Name: "arch", Value: "aarch64"},
					{Name: "os", Value: "linux"},
				},
			},
		}, key.GetPlatformQueueName())
	})

	t.Run("NoPlatformPresent", func(t *testing.T) {
		// If no platform object is, assume the empty set of
		// platform properties. Clients such as BuildStream are
		// known for not providing them.
		key, err := keyExtractor.ExtractKey(ctx, instanceName, &remoteexecution.Action{})
		require.NoError(t, err)
		testutil.RequireEqualProto(t, &buildqueuestate.PlatformQueueName{
			InstanceNamePrefix: "hello",
			Platform:           &remoteexecution.Platform{},
		}, key.GetPlatformQueueName())
	})
}
