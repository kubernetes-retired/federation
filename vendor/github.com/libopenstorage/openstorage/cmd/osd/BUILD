load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/libopenstorage/openstorage/cmd/osd",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/codegangsta/cli:go_default_library",
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/flexvolume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/server:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/cli:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/cluster:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/config:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/github.com/portworx/kvdb/consul:go_default_library",
        "//vendor/github.com/portworx/kvdb/etcd/v2:go_default_library",
        "//vendor/github.com/portworx/kvdb/mem:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
    ],
)

go_binary(
    name = "osd",
    importpath = "github.com/libopenstorage/openstorage/cmd/osd",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
