load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["drivers.go"],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/aws:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/btrfs:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/buse:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/coprhd:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/nfs:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/pwx:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/vfs:go_default_library",
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
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/aws:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/btrfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/buse:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/common:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/coprhd:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/fuse:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/nfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/pwx:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/test:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/vfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
