load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["templates.go"],
    importpath = "github.com/docker/docker/pkg/templates",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["templates_test.go"],
    importpath = "github.com/docker/docker/pkg/templates",
    library = ":go_default_library",
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
