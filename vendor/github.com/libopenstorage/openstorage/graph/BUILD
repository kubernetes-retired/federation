load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["graphdriver.go"],
    importpath = "github.com/libopenstorage/openstorage/graph",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/graphdriver:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
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
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
