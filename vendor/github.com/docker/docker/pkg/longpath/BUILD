load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["longpath.go"],
    importpath = "github.com/docker/docker/pkg/longpath",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["longpath_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/docker/docker/pkg/longpath",
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
