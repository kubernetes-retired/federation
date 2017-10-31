load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["auth.go"],
    importpath = "github.com/docker/distribution/registry/auth",
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
        "//vendor/github.com/docker/distribution/registry/auth/htpasswd:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth/silly:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth/token:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
