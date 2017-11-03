load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "default_store_enumerator.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/proto/time:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["default_store_enumerator_test.go"],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/common",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/github.com/portworx/kvdb/mem:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
