load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fileinfo.go",
        "storagedriver.go",
    ],
    importpath = "github.com/docker/distribution/registry/storage/driver",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/distribution/context:go_default_library"],
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
        "//vendor/github.com/docker/distribution/registry/storage/driver/azure:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/base:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/factory:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/filesystem:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/gcs:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/inmemory:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/middleware:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/oss:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/s3-aws:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/s3-goamz:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/swift:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/testdriver:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/testsuites:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
