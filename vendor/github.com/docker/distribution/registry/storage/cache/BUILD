load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cache.go",
        "cachedblobdescriptorstore.go",
    ],
    importpath = "github.com/docker/distribution/registry/storage/cache",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/cache/cachecheck:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/cache/memory:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/cache/redis:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
