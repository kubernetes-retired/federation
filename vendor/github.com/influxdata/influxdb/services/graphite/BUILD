load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "errors.go",
        "parser.go",
        "service.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/graphite",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/monitor/diagnostics:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
        "//vendor/github.com/influxdata/influxdb/toml:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["service_test.go"],
    importpath = "github.com/influxdata/influxdb/services/graphite",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/influxdata/influxdb/internal:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
        "//vendor/github.com/influxdata/influxdb/toml:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "config_test.go",
        "parser_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/graphite_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
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
