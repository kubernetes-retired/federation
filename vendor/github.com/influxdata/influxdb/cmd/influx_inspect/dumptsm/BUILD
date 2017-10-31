load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["dumptsm.go"],
    importpath = "github.com/influxdata/influxdb/cmd/influx_inspect/dumptsm",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/influxdata/influxdb/tsdb/engine/tsm1:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["dumptsm_test.go"],
    importpath = "github.com/influxdata/influxdb/cmd/influx_inspect/dumptsm_test",
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
