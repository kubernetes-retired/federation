load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "tls_config.go",
    ],
    importpath = "github.com/prometheus/common/config",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["tls_config_test.go"],
    data = glob(["testdata/**"]),
    importpath = "github.com/prometheus/common/config",
    library = ":go_default_library",
    deps = ["//vendor/gopkg.in/yaml.v2:go_default_library"],
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
