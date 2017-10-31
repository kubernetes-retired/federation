load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["gen.go"],
    importpath = "github.com/Azure/azure-sdk-for-go/Gododir",
    visibility = ["//visibility:private"],
    deps = ["//vendor/gopkg.in/godo.v2:go_default_library"],
)

go_binary(
    name = "Gododir",
    importpath = "github.com/Azure/azure-sdk-for-go/Gododir",
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
