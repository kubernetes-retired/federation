load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["run.go"],
    importpath = "github.com/coreos/rkt/stage1/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
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
        "//vendor/github.com/coreos/rkt/stage1/common/ssh:all-srcs",
        "//vendor/github.com/coreos/rkt/stage1/common/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
