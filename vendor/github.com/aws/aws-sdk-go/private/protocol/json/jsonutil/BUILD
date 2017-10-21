load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "build.go",
        "unmarshal.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/aws/aws-sdk-go/private/protocol:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["build_test.go"],
    importpath = "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
