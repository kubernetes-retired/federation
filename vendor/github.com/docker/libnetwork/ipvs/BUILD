load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "constants.go",
            "ipvs.go",
            "netlink.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/ipvs",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/sirupsen/logrus:go_default_library",
            "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
            "//vendor/github.com/vishvananda/netns:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "ipvs_test.go",
        ],
        "//conditions:default": [],
    }),
    embed = [":go_default_library"],
    importpath = "github.com/docker/libnetwork/ipvs",
    deps = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
            "//vendor/github.com/stretchr/testify/assert:go_default_library",
            "//vendor/github.com/stretchr/testify/require:go_default_library",
            "//vendor/github.com/vishvananda/netlink:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
