load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["image_reference.go"],
    importpath = "github.com/docker/engine-api/types/reference",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/distribution/reference:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["image_reference_test.go"],
    importpath = "github.com/docker/engine-api/types/reference",
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
