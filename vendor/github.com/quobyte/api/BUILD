load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "quobyte.go",
        "rpc_client.go",
        "types.go",
    ],
    importpath = "github.com/quobyte/api",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["rpc_client_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/quobyte/api",
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
