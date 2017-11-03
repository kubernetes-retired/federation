load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["layer0.go"],
    importpath = "github.com/libopenstorage/openstorage/graph/drivers/layer0",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/graphdriver:go_default_library",
        "//vendor/github.com/docker/docker/daemon/graphdriver/overlay:go_default_library",
        "//vendor/github.com/docker/docker/pkg/archive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/docker/docker/pkg/parsers:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/graph:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
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
