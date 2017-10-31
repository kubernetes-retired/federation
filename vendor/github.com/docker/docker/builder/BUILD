load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["builder.go"],
    importpath = "github.com/docker/docker/builder",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/backend:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/container:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
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
        "//vendor/github.com/docker/docker/builder/dockerfile:all-srcs",
        "//vendor/github.com/docker/docker/builder/dockerignore:all-srcs",
        "//vendor/github.com/docker/docker/builder/fscache:all-srcs",
        "//vendor/github.com/docker/docker/builder/remotecontext:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
