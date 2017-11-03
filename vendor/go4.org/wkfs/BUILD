load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["wkfs.go"],
    importpath = "go4.org/wkfs",
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
    srcs = [
        ":package-srcs",
        "//vendor/go4.org/wkfs/gcs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
