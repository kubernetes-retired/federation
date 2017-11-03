load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["storagemiddleware.go"],
    importpath = "github.com/docker/distribution/registry/storage/driver/middleware",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/distribution/registry/storage/driver:go_default_library"],
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
        "//vendor/github.com/docker/distribution/registry/storage/driver/middleware/cloudfront:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/middleware/redirect:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
