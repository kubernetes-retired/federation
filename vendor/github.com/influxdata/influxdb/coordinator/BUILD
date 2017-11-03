load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "meta_client.go",
        "points_writer.go",
        "statement_executor.go",
    ],
    importpath = "github.com/influxdata/influxdb/coordinator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/influxdata/influxdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/monitor:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
        "//vendor/github.com/influxdata/influxdb/toml:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["points_writer_internal_test.go"],
    importpath = "github.com/influxdata/influxdb/coordinator",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "config_test.go",
        "meta_client_test.go",
        "points_writer_test.go",
        "statement_executor_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/coordinator_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/influxdata/influxdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/internal:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
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
