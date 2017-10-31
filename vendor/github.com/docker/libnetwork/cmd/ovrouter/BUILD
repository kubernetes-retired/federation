load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["ovrouter.go"],
    importpath = "github.com/docker/libnetwork/cmd/ovrouter",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/drivers/overlay:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
    ],
)

go_binary(
    name = "ovrouter",
    importpath = "github.com/docker/libnetwork/cmd/ovrouter",
    library = ":go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
