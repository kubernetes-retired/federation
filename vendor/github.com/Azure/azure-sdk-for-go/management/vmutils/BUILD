load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "configurationset.go",
        "datadisks.go",
        "deployment.go",
        "extensions.go",
        "network.go",
        "rolesize.go",
        "rolestate.go",
        "vmutils.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/management/vmutils",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/location:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachine:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "extensions_test.go",
        "integration_test.go",
        "vmutils_test.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/management/vmutils",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/hostedservice:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/location:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/osimage:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/storageservice:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/testutils:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachine:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachineimage:go_default_library",
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
