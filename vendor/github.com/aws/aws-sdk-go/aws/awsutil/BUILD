load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "copy.go",
        "equal.go",
        "path_value.go",
        "prettify.go",
        "string_value.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/aws/awsutil",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/jmespath/go-jmespath:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "copy_test.go",
        "equal_test.go",
        "path_value_test.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/aws/awsutil_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
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
