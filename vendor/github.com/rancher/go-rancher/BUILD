load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/rancher/go-rancher",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/rancher/go-rancher/client:go_default_library"],
)

go_binary(
    name = "go-rancher",
    importpath = "github.com/rancher/go-rancher",
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
        "//vendor/github.com/rancher/go-rancher/api:all-srcs",
        "//vendor/github.com/rancher/go-rancher/catalog:all-srcs",
        "//vendor/github.com/rancher/go-rancher/client:all-srcs",
        "//vendor/github.com/rancher/go-rancher/docs:all-srcs",
        "//vendor/github.com/rancher/go-rancher/generator:all-srcs",
        "//vendor/github.com/rancher/go-rancher/hostaccess:all-srcs",
        "//vendor/github.com/rancher/go-rancher/v2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
