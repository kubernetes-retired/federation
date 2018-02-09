load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cast.go",
        "caste.go",
    ],
    importpath = "github.com/spf13/cast",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/spf13/jwalterweatherman:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["cast_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/spf13/cast",
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
