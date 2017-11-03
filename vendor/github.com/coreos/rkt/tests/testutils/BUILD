load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ctx.go",
        "datadirsetup.go",
        "goroutineassistant.go",
        "httputils.go",
        "iputils.go",
        "test.go",
        "utils.go",
    ],
    importpath = "github.com/coreos/rkt/tests/testutils",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/gexpect:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/group:go_default_library",
        "//vendor/github.com/coreos/rkt/tests/testutils/logger:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/github.com/hydrogen18/stoppableListener:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
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
        "//vendor/github.com/coreos/rkt/tests/testutils/aci-server:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/testutils/logger:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
