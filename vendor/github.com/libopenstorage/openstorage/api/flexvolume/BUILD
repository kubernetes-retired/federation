load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["flexvolume.go"],
    importpath = "github.com/libopenstorage/openstorage/api/flexvolume",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/pkg/flexvolume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/mount:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
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
