load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["nfs.go"],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/nfs",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/config:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/mount:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/seed:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/common:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["nfs_test.go"],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/nfs",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/test:go_default_library",
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
