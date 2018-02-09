load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "join.go",
        "vfs.go",
    ],
    importpath = "github.com/cyphar/filepath-securejoin",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pkg/errors:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["join_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/cyphar/filepath-securejoin",
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
