load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["resumablerequestreader.go"],
    importpath = "github.com/docker/docker/registry/resumable",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/sirupsen/logrus:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["resumablerequestreader_test.go"],
    importpath = "github.com/docker/docker/registry/resumable",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
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
