load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["create.go"],
    importpath = "github.com/Azure/azure-sdk-for-go/arm/examples/dns",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/dns:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/examples/helpers:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
    ],
)

go_binary(
    name = "dns",
    importpath = "github.com/Azure/azure-sdk-for-go/arm/examples/dns",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/examples/dns/paging:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
