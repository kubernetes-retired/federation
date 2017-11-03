load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cluster.go",
        "database.go",
        "manager.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/cluster",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/gossip:go_default_library",
        "//vendor/github.com/libopenstorage/gossip/types:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/config:go_default_library",
        "//vendor/github.com/libopenstorage/systemutils:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
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
