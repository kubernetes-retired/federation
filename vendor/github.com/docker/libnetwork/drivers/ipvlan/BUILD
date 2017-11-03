load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "ipvlan.go",
        "ipvlan_endpoint.go",
        "ipvlan_joinleave.go",
        "ipvlan_network.go",
        "ipvlan_setup.go",
        "ipvlan_state.go",
        "ipvlan_store.go",
    ],
    importpath = "github.com/docker/libnetwork/drivers/ipvlan",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/parsers/kernel:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/ns:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/osl:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "ipvlan_setup_test.go",
        "ipvlan_test.go",
    ],
    importpath = "github.com/docker/libnetwork/drivers/ipvlan",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
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
        "//vendor/github.com/docker/libnetwork/drivers/ipvlan/ivmanager:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
