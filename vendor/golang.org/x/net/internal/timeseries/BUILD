load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["timeseries.go"],
    importpath = "golang.org/x/net/internal/timeseries",
    visibility = ["//vendor/golang.org/x/net:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["timeseries_test.go"],
    embed = [":go_default_library"],
    importpath = "golang.org/x/net/internal/timeseries",
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
