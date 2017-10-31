load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "errors.go",
        "http.go",
        "operations.go",
        "publishSettings.go",
        "util.go",
        "version.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/management",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/pkcs12:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["errors_test.go"],
    importpath = "github.com/Azure/azure-sdk-for-go/management_test",
    deps = [":go_default_library"],
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
        "//vendor/github.com/Azure/azure-sdk-for-go/management/affinitygroup:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/hostedservice:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/location:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/networksecuritygroup:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/osimage:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/sql:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/storageservice:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/testutils:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachine:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachineimage:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualnetwork:all-srcs",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/vmutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
