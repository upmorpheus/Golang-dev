load("@bazel_gazelle//:deps.bzl", "go_repository")

def bb_storage_go_dependencies():
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sha256 = "6ba3169493880a63128b6c6edc9119817df257db0b34b27887cad871767f0525",
        strip_prefix = "aws-sdk-go-1.16.26",
        urls = ["https://github.com/aws/aws-sdk-go/archive/v1.16.26.tar.gz"],
    )

    go_repository(
        name = "com_github_bazelbuild_remote_apis",
        importpath = "github.com/bazelbuild/remote-apis",
        patches = [
            "@com_github_buildbarn_bb_storage//:patches/com_github_bazelbuild_remote_apis/auxiliary_metadata.diff",
            "@com_github_buildbarn_bb_storage//:patches/com_github_bazelbuild_remote_apis/golang.diff",
        ],
        sha256 = "6fee1d4d911324cd0a95aa19411867a17bb8c34a9a90667c21c825b122b0d79a",
        strip_prefix = "remote-apis-7802003e00901b4e740fe0ebec1243c221e02ae2",
        urls = ["https://github.com/bazelbuild/remote-apis/archive/7802003e00901b4e740fe0ebec1243c221e02ae2.tar.gz"],
    )

    go_repository(
        name = "com_github_beorn7_perks",
        commit = "3a771d992973f24aa725d07868b467d1ddfceafb",
        importpath = "github.com/beorn7/perks",
    )

    go_repository(
        name = "com_github_buildbarn_bb_storage",
        commit = "dfb8c06f0dda1d945284616c75ed4b3706906b8b",
        importpath = "github.com/buildbarn/bb-storage",
    )

    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        patches = ["@com_github_buildbarn_bb_storage//:patches/com_github_golang_mock/mocks-for-funcs.diff"],
        sha256 = "d6a80e2ba2d23b5cbcce09eaff6d8a961262f12b95ae18011ee9ad4ddd339f76",
        strip_prefix = "mock-6d816de489c18a7e9a8fbd2aa5bb2dd75f2bbc86",
        urls = ["https://github.com/golang/mock/archive/6d816de489c18a7e9a8fbd2aa5bb2dd75f2bbc86.tar.gz"],
    )

    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sha256 = "7e330758f7c81d9f489493fb7ae0e67d06f50753429758b64f25ad5fb2727e21",
        strip_prefix = "uuid-1.1.0",
        urls = ["https://github.com/google/uuid/archive/v1.1.0.tar.gz"],
    )

    go_repository(
        name = "com_github_go_redis_redis",
        importpath = "github.com/go-redis/redis",
        sha256 = "c997aca07026a52745e3d7aeab528550b90d3bae65ff2b67991d890bb2a7b1ff",
        strip_prefix = "redis-6.15.1",
        urls = ["https://github.com/go-redis/redis/archive/v6.15.1.tar.gz"],
    )

    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sha256 = "eba66530952a126ab869205bdb909af607bfd9eb09f00207b62eb29140258aa9",
        strip_prefix = "go-grpc-prometheus-1.2.0",
        urls = ["https://github.com/grpc-ecosystem/go-grpc-prometheus/archive/v1.2.0.tar.gz"],
    )

    go_repository(
        name = "com_github_lazybeaver_xorshift",
        commit = "ce511d4823dd074d7c37a74225320332d6961abb",
        importpath = "github.com/lazybeaver/xorshift",
    )

    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        commit = "c12348ce28de40eed0136aa2b644d0ee0650e56c",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
    )

    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sha256 = "e255f632b7223f794e0f60d99976535c16153cc00fadf39ee87fd9f678d6a32c",
        strip_prefix = "client_golang-1.6.0",
        urls = ["https://github.com/prometheus/client_golang/archive/v1.6.0.tar.gz"],
    )

    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sha256 = "0ee31178d2c5a1249be4e26294a2f428008dc4e1ecbbfbe47f74e41026df1148",
        strip_prefix = "xxhash-2.1.1",
        urls = ["https://github.com/cespare/xxhash/archive/v2.1.1.tar.gz"],
    )

    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sha256 = "4ab1be9cdfa702d7f49beeb09a256bcc6a2aad55e8a0a37e7732a46934264e12",
        strip_prefix = "client_model-0.2.0",
        urls = ["https://github.com/prometheus/client_model/archive/v0.2.0.tar.gz"],
    )

    go_repository(
        name = "com_github_prometheus_common",
        commit = "6225330c4d5d7d0ffa21a82ad7eb041f82f25c7f",
        importpath = "github.com/prometheus/common",
    )

    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sha256 = "7f31adcafbbfd7cb2e5d5b5954a3f55565ad41fb039c977b1f769c7d3dfc2a60",
        strip_prefix = "procfs-0.0.10",
        urls = ["https://github.com/prometheus/procfs/archive/v0.0.10.tar.gz"],
    )

    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sha256 = "3ae072321569a8cd6d77de8f3be774165e136198ce808df0a31589237ba59698",
        strip_prefix = "testify-1.4.0",
        urls = ["https://github.com/stretchr/testify/archive/v1.4.0.tar.gz"],
    )

    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sha256 = "74afe4be0fa4482de73cfaf952dc6c8d41088b687c3f9de4a7ced86d2bbe0cb4",
        strip_prefix = "yaml-2.2.7",
        urls = ["https://github.com/go-yaml/yaml/archive/v2.2.7.tar.gz"],
    )

    go_repository(
        name = "io_opencensus_go_contrib_exporter_prometheus",
        importpath = "contrib.go.opencensus.io/exporter/prometheus",
        commit = "f6cda26f80a388eabda7766388c14e96370440e5",
    )

    go_repository(
        name = "io_opencensus_go_contrib_exporter_jaeger",
        commit = "e8b55949d948652e47aae4378212f933ecee856b",
        importpath = "contrib.go.opencensus.io/exporter/jaeger",
    )

    go_repository(
        name = "dev_gocloud",
        importpath = "gocloud.dev",
        sha256 = "e8d952e0a78473e822b368c5ee1d6a45d2e531d76758db86ee56dbd1eb623f6e",
        strip_prefix = "go-cloud-0.18.0",
        urls = ["https://github.com/google/go-cloud/archive/v0.18.0.tar.gz"],
    )

    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        urls = ["https://github.com/googleapis/google-api-go-client/archive/v0.4.0.tar.gz"],
        sha256 = "fde7b06bc002cc886efa94845ac2ba4f48fd4c321a04a9ee5558026f5fa28c0c",
        strip_prefix = "google-api-go-client-0.4.0",
    )

    go_repository(
        name = "com_github_uber_jaeger_client_go",
        importpath = "github.com/uber/jaeger-client-go",
        urls = ["https://github.com/jaegertracing/jaeger-client-go/archive/v2.16.0.tar.gz"],
        sha256 = "9657eb6603d6aae55c5637957ab63400127bcc395981831366998428cc3f7edb",
        strip_prefix = "jaeger-client-go-2.16.0",
    )

    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        commit = "112230192c580c3556b8cee6403af37a4fc5f28c",
    )

    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        urls = ["https://github.com/census-instrumentation/opencensus-go/archive/v0.21.0.tar.gz"],
        sha256 = "e7129aebb9bcb590f01b4fb773b6cf0b10109211cb38cfbaf1f097d191043251",
        strip_prefix = "opencensus-go-0.21.0",
    )

    go_repository(
        name = "com_google_cloud_go",
        commit = "09ad026a62f0561b7f7e276569eda11a6afc9773",
        importpath = "cloud.google.com/go",
    )

    go_repository(
        name = "org_golang_x_xerrors",
        commit = "a985d3407aa71f30cf86696ee0a2f409709f22e1",
        importpath = "golang.org/x/xerrors",
    )

    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        urls = ["https://github.com/hashicorp/golang-lru/archive/v0.5.1.tar.gz"],
        sha256 = "3bf57512af746dc0338651ba1c35c65fe907ff214ccb22d679539f7ea791511e",
        strip_prefix = "golang-lru-0.5.1",
    )

    go_repository(
        name = "com_github_googleapis_gax_go",
        importpath = "github.com/googleapis/gax-go",
        sha256 = "3089affe6f5e27f7a6d494cb399aa6baf232384f763f548ad5ddfbea0e88e59c",
        strip_prefix = "gax-go-2.0.5",
        urls = ["https://github.com/googleapis/gax-go/archive/v2.0.5.tar.gz"],
    )

    go_repository(
        name = "org_golang_x_oauth2",
        commit = "9f3314589c9a9136388751d9adae6b0ed400978a",
        importpath = "golang.org/x/oauth2",
    )

    go_repository(
        name = "com_github_google_wire",
        commit = "2183ee4806cf1878e136fea26f06f9abef9375b6",
        importpath = "github.com/google/wire",
        build_extra_args = ["--exclude=internal/wire/testdata"],
    )

    go_repository(
        name = "com_github_azure_azure_pipeline_go",
        importpath = "github.com/Azure/azure-pipeline-go",
        sha256 = "dc0d15949088e17e74da35c3ae2730f52240ad73e25cf795f532b7282de68e2f",
        strip_prefix = "azure-pipeline-go-0.2.2",
        urls = ["https://github.com/Azure/azure-pipeline-go/archive/v0.2.2.tar.gz"],
    )

    go_repository(
        name = "com_github_azure_azure_storage_blob_go",
        importpath = "github.com/Azure/azure-storage-blob-go",
        sha256 = "e50db03bc05b952159e0a1d393c841c217c3aecd6ce745f8b9c42fd498e476c7",
        strip_prefix = "azure-storage-blob-go-0.8.0",
        urls = ["https://github.com/Azure/azure-storage-blob-go/archive/v0.8.0.tar.gz"],
    )

    go_repository(
        name = "com_github_google_go_jsonnet",
        commit = "0959f85501584da690e34871b31e280242226e1f",
        importpath = "github.com/google/go-jsonnet",
        patches = ["@com_github_buildbarn_bb_storage//:patches/com_github_google_go_jsonnet/astgen.diff"],
    )

    go_repository(
        name = "com_github_fatih_color",
        commit = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
        importpath = "github.com/fatih/color",
    )

    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        commit = "cfaf5686ec79ff8344257723b6f5ba1ae0ffeb4d",
    )

    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        commit = "49c01487a141b49f8ffe06277f3dca3ee80a55fa",
    )

    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        importpath = "github.com/googleapis/gax-go/v2",
        sha256 = "3089affe6f5e27f7a6d494cb399aa6baf232384f763f548ad5ddfbea0e88e59c",
        strip_prefix = "gax-go-2.0.5/v2",
        urls = ["https://github.com/googleapis/gax-go/archive/v2.0.5.tar.gz"],
    )

    go_repository(
        name = "com_github_mattn_go_ieproxy",
        commit = "7c0f6868bffe087073376feaab3ace57f2ef90b2",
        importpath = "github.com/mattn/go-ieproxy",
    )

    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        commit = "6af3d372ceca1a2c17f8c7789446a3488a91b8c6",
        importpath = "google.golang.org/grpc",
    )

    go_repository(
        name = "com_github_gordonklaus_ineffassign",
        commit = "7953dde2c7bf4ce700d9f14c2e41c0966763760c",
        importpath = "github.com/gordonklaus/ineffassign",
    )

    go_repository(
        name = "org_golang_x_lint",
        commit = "738671d3881b9731cc63024d5d88cf28db875626",
        importpath = "golang.org/x/lint",
    )
