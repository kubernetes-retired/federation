load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["names-generator.go"],
    importpath = "github.com/docker/docker/pkg/namesgenerator",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["names-generator_test.go"],
    importpath = "github.com/docker/docker/pkg/namesgenerator",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/docker/pkg/namesgenerator/cmd/names-generator:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
