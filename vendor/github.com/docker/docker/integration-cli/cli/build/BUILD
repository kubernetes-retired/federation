load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["build.go"],
    importpath = "github.com/docker/docker/integration-cli/cli/build",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakecontext:go_default_library",
        "//vendor/github.com/docker/docker/pkg/testutil/cmd:go_default_library",
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
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakecontext:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakegit:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakestorage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
