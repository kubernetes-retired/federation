load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "addr.go",
        "class.go",
        "conntrack_unspecified.go",
        "filter.go",
        "genetlink_unspecified.go",
        "handle_unspecified.go",
        "link.go",
        "neigh.go",
        "netlink.go",
        "netlink_unspecified.go",
        "order.go",
        "protinfo.go",
        "qdisc.go",
        "route.go",
        "route_unspecified.go",
        "rule.go",
        "socket.go",
        "xfrm.go",
        "xfrm_policy.go",
        "xfrm_state.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "addr_linux.go",
            "bpf_linux.go",
            "bridge_linux.go",
            "class_linux.go",
            "conntrack_linux.go",
            "filter_linux.go",
            "genetlink_linux.go",
            "gtp_linux.go",
            "handle_linux.go",
            "link_linux.go",
            "link_tuntap_linux.go",
            "neigh_linux.go",
            "netlink_linux.go",
            "protinfo_linux.go",
            "qdisc_linux.go",
            "route_linux.go",
            "rule_linux.go",
            "socket_linux.go",
            "xfrm_monitor_linux.go",
            "xfrm_policy_linux.go",
            "xfrm_state_linux.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    importpath = "github.com/vishvananda/netlink",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
        "//vendor/github.com/vishvananda/netns:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "conntrack_test.go",
        "netlink_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "addr_test.go",
            "bridge_linux_test.go",
            "class_test.go",
            "filter_test.go",
            "gtp_test.go",
            "handle_test.go",
            "link_test.go",
            "neigh_test.go",
            "protinfo_test.go",
            "qdisc_test.go",
            "route_test.go",
            "rule_test.go",
            "socket_test.go",
            "xfrm_monitor_test.go",
            "xfrm_policy_test.go",
            "xfrm_state_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/vishvananda/netlink",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/vishvananda/netns:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
        "//vendor/github.com/vishvananda/netlink/nl:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
