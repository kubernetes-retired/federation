load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["tailfile.go"],
    importpath = "github.com/docker/docker/pkg/tailfile",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["tailfile_test.go"],
    importpath = "github.com/docker/docker/pkg/tailfile",
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
