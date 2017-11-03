load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["drivers.go"],
    importpath = "github.com/libopenstorage/openstorage/graph/drivers",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/chainfs:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/layer0:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/proxy:go_default_library",
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
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/chainfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/layer0:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/proxy:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
