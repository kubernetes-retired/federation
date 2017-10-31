load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/influxdata/influxdb/services/collectd/test_client",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/collectd.org/api:go_default_library",
        "//vendor/collectd.org/network:go_default_library",
    ],
)

go_binary(
    name = "test_client",
    importpath = "github.com/influxdata/influxdb/services/collectd/test_client",
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
