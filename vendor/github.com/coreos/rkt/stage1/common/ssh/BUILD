load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["ssh.go"],
    importpath = "github.com/coreos/rkt/stage1/common/ssh",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/networking/netinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
