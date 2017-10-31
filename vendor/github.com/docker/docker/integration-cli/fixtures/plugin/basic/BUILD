load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["basic.go"],
    importpath = "github.com/docker/docker/integration-cli/fixtures/plugin/basic",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "basic",
    importpath = "github.com/docker/docker/integration-cli/fixtures/plugin/basic",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
