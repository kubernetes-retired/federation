load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["influxdb.go"],
    importpath = "github.com/google/cadvisor/storage/influxdb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/google/cadvisor/storage:go_default_library",
        "//vendor/github.com/google/cadvisor/version:go_default_library",
        "//vendor/github.com/influxdb/influxdb/client:go_default_library",
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
