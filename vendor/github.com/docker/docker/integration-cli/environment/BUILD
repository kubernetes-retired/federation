load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "clean.go",
        "environment.go",
        "protect.go",
    ],
    importpath = "github.com/docker/docker/integration-cli/environment",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/api/types/volume:go_default_library",
        "//vendor/github.com/docker/docker/client:go_default_library",
        "//vendor/github.com/docker/docker/integration-cli/request:go_default_library",
        "//vendor/github.com/docker/docker/opts:go_default_library",
        "//vendor/github.com/docker/docker/pkg/testutil/cmd:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
