load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["check.go"],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/examples/check",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/examples/helpers:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/storage:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
    ],
)

go_binary(
    name = "check",
    importpath = "github.com/Azure/azure-sdk-for-go/arm/examples/check",
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
