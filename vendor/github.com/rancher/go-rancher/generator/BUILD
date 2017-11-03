load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["generator.go"],
    importpath = "github.com/rancher/go-rancher/generator",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/rancher/go-rancher/client:go_default_library"],
)

go_binary(
    name = "generator",
    importpath = "github.com/rancher/go-rancher/generator",
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
