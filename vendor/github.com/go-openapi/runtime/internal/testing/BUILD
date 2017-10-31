load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["data.go"],
    importpath = "github.com/go-openapi/runtime/internal/testing",
    visibility = ["//vendor/github.com/go-openapi/runtime:__subpackages__"],
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
        "//vendor/github.com/go-openapi/runtime/internal/testing/petstore:all-srcs",
        "//vendor/github.com/go-openapi/runtime/internal/testing/simplepetstore:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
