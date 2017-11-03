load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["parser.go"],
    importpath = "github.com/influxdata/influxdb/stress/v2/stressql/statement",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/influxdata/influxdb/stress/v2/statement:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["parser_test.go"],
    importpath = "github.com/influxdata/influxdb/stress/v2/stressql/statement",
    library = ":go_default_library",
    deps = ["//vendor/github.com/influxdata/influxdb/stress/v2/statement:go_default_library"],
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
