load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["ovmanager.go"],
    importpath = "github.com/docker/libnetwork/drivers/solaris/overlay/ovmanager",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/idm:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
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
