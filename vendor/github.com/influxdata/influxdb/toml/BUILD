load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["toml.go"],
    importpath = "github.com/influxdata/influxdb/toml",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["toml_test.go"],
    importpath = "github.com/influxdata/influxdb/toml_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/influxdata/influxdb/cmd/influxd/run:go_default_library",
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
