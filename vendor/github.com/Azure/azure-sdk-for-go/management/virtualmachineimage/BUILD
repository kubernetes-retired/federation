load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "entities.go",
    ],
    importpath = "github.com/Azure/azure-sdk-for-go/management/virtualmachineimage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["entities_test.go"],
    importpath = "github.com/Azure/azure-sdk-for-go/management/virtualmachineimage",
    library = ":go_default_library",
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
