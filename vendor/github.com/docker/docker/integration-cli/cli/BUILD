load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cli.go"],
    importpath = "github.com/docker/docker/integration-cli/cli",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/integration-cli/daemon:go_default_library",
        "//vendor/github.com/docker/docker/integration-cli/environment:go_default_library",
        "//vendor/github.com/docker/docker/pkg/testutil/cmd:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli/build:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
