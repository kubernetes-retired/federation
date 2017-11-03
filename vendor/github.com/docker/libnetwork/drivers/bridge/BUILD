load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bridge.go",
        "bridge_store.go",
        "errors.go",
        "interface.go",
        "labels.go",
        "link.go",
        "netlink_deprecated_linux_notarm.go",
        "netlink_deprecated_unsupported.go",
        "port_mapping.go",
        "setup.go",
        "setup_bridgenetfiltering.go",
        "setup_device.go",
        "setup_firewalld.go",
        "setup_ip_forwarding.go",
        "setup_ip_tables.go",
        "setup_ipv4.go",
        "setup_ipv6.go",
        "setup_verify.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "netlink_deprecated_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/drivers/bridge",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/parsers/kernel:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/iptables:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/ns:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/osl:go_default_library",
        "//vendor/github.com/docker/libnetwork/portmapper:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bridge_test.go",
        "interface_test.go",
        "link_test.go",
        "network_test.go",
        "port_mapping_test.go",
        "setup_bridgenetfiltering_test.go",
        "setup_device_test.go",
        "setup_ip_forwarding_test.go",
        "setup_ip_tables_test.go",
        "setup_ipv4_test.go",
        "setup_ipv6_test.go",
        "setup_verify_test.go",
    ],
    importpath = "github.com/docker/libnetwork/drivers/bridge",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/iptables:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/portmapper:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
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
        "//vendor/github.com/docker/libnetwork/drivers/bridge/brmanager:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
