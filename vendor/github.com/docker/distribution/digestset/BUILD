load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["set.go"],
    importpath = "github.com/docker/distribution/digestset",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/go-digest:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["set_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/docker/distribution/digestset",
    deps = ["//vendor/github.com/opencontainers/go-digest:go_default_library"],
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
