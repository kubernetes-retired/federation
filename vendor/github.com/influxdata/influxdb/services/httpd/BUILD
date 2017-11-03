load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "handler.go",
        "listen.go",
        "response_logger.go",
        "response_writer.go",
        "service.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/httpd",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/bmizerany/pat:go_default_library",
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/influxdata/influxdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/monitor:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/uuid:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "config_test.go",
        "handler_test.go",
        "listen_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/httpd_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta:go_default_library",
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
