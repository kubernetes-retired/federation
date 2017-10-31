load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["instance.go"],
    importpath = "github.com/rancher/go-rancher/hostaccess",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gorilla/websocket:go_default_library",
        "//vendor/github.com/rancher/go-rancher/v2:go_default_library",
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
