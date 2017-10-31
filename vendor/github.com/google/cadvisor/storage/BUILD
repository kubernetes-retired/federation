load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "common_flags.go",
        "storage.go",
    ],
    importpath = "github.com/google/cadvisor/storage",
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
        "//vendor/github.com/google/cadvisor/storage/bigquery:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/elasticsearch:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/influxdb:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/kafka:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/redis:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/statsd:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/stdout:all-srcs",
        "//vendor/github.com/google/cadvisor/storage/test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
