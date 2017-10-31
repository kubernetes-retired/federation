load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["gpio.go"],
    importpath = "golang.org/x/exp/io/gpio",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/exp/io/gpio/driver:go_default_library"],
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
        "//vendor/golang.org/x/exp/io/gpio/driver:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
