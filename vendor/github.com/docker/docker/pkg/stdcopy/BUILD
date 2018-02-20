load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["stdcopy.go"],
    importpath = "github.com/docker/docker/pkg/stdcopy",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["stdcopy_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/docker/docker/pkg/stdcopy",
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
