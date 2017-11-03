load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["pools.go"],
    importpath = "github.com/docker/docker/pkg/pools",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/ioutils:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["pools_test.go"],
    importpath = "github.com/docker/docker/pkg/pools",
    library = ":go_default_library",
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
