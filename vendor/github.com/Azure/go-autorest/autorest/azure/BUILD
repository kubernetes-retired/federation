load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "async.go",
        "azure.go",
        "environments.go",
    ],
    importpath = "github.com/Azure/go-autorest/autorest/azure",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "async_test.go",
        "azure_test.go",
        "environments_test.go",
    ],
    importpath = "github.com/Azure/go-autorest/autorest/azure",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/mocks:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/Azure/go-autorest/autorest/azure/example:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
