load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["checker.go"],
    importpath = "github.com/docker/docker/integration-cli/checker",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/go-check/check:go_default_library",
        "//vendor/github.com/vdemeester/shakers:go_default_library",
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
