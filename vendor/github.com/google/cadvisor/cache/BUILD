load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cache.go"],
    importpath = "github.com/google/cadvisor/cache",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/google/cadvisor/info/v1:go_default_library"],
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
        "//vendor/github.com/google/cadvisor/cache/memory:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
