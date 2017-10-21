load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "nl_unspecified.go",
        "syscall.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "addr_linux.go",
            "bridge_linux.go",
            "conntrack_linux.go",
            "genetlink_linux.go",
            "link_linux.go",
            "mpls_linux.go",
            "nl_linux.go",
            "route_linux.go",
            "tc_linux.go",
            "xfrm_linux.go",
            "xfrm_monitor_linux.go",
            "xfrm_policy_linux.go",
            "xfrm_state_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/vishvananda/netlink/nl",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/vishvananda/netns:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "addr_linux_test.go",
            "bridge_linux_test.go",
            "link_linux_test.go",
            "nl_linux_test.go",
            "route_linux_test.go",
            "tc_linux_test.go",
            "xfrm_linux_test.go",
            "xfrm_monitor_linux_test.go",
            "xfrm_policy_linux_test.go",
            "xfrm_state_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/vishvananda/netlink/nl",
    library = ":go_default_library",
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
