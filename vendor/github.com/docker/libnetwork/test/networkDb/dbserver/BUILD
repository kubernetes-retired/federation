load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["ndbServer.go"],
    importpath = "github.com/docker/libnetwork/test/networkDb/dbserver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/diagnose:go_default_library",
        "//vendor/github.com/docker/libnetwork/networkdb:go_default_library",
        "//vendor/github.com/docker/libnetwork/test/networkDb/dummyclient:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
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
