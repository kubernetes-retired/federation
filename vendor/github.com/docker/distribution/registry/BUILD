load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "registry.go",
        "root.go",
    ],
    importpath = "github.com/docker/distribution/registry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bshuster-repo/logrus-logstash-hook:go_default_library",
        "//vendor/github.com/bugsnag/bugsnag-go:go_default_library",
        "//vendor/github.com/docker/distribution/configuration:go_default_library",
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/health:go_default_library",
        "//vendor/github.com/docker/distribution/registry/handlers:go_default_library",
        "//vendor/github.com/docker/distribution/registry/listener:go_default_library",
        "//vendor/github.com/docker/distribution/registry/storage:go_default_library",
        "//vendor/github.com/docker/distribution/registry/storage/driver/factory:go_default_library",
        "//vendor/github.com/docker/distribution/uuid:go_default_library",
        "//vendor/github.com/docker/distribution/version:go_default_library",
        "//vendor/github.com/docker/libtrust:go_default_library",
        "//vendor/github.com/gorilla/handlers:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/yvasiyarov/gorelic:go_default_library",
        "//vendor/rsc.io/letsencrypt:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["registry_test.go"],
    importpath = "github.com/docker/distribution/registry",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/distribution/configuration:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/distribution/registry/api/errcode:all-srcs",
        "//vendor/github.com/docker/distribution/registry/api/v2:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth:all-srcs",
        "//vendor/github.com/docker/distribution/registry/client:all-srcs",
        "//vendor/github.com/docker/distribution/registry/handlers:all-srcs",
        "//vendor/github.com/docker/distribution/registry/listener:all-srcs",
        "//vendor/github.com/docker/distribution/registry/middleware/registry:all-srcs",
        "//vendor/github.com/docker/distribution/registry/middleware/repository:all-srcs",
        "//vendor/github.com/docker/distribution/registry/proxy:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
