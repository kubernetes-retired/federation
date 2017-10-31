load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["help.go"],
    importpath = "github.com/influxdata/influxdb/cmd/influx_inspect/help",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["help_test.go"],
    importpath = "github.com/influxdata/influxdb/cmd/influx_inspect/help_test",
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
