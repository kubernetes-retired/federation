load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "conntrack.go",
        "firewalld.go",
        "iptables.go",
    ],
    importpath = "github.com/docker/libnetwork/iptables",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/godbus/dbus:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "firewalld_test.go",
        "iptables_test.go",
    ],
    importpath = "github.com/docker/libnetwork/iptables",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/libnetwork/testutils:go_default_library"],
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
