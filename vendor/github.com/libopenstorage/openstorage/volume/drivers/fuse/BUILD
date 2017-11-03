load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fuse.go",
        "volume_driver.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/fuse",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/bazil.org/fuse:go_default_library",
        "//vendor/bazil.org/fuse/fs:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/common:go_default_library",
        "//vendor/github.com/pborman/uuid:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
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
