load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["common.go"],
    importpath = "github.com/appc/spec/schema/common",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["common_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/appc/spec/schema/common",
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
