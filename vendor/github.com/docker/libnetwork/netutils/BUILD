load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "utils.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "utils_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "utils_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/netutils",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
            "//vendor/github.com/docker/libnetwork/ns:go_default_library",
            "//vendor/github.com/docker/libnetwork/osl:go_default_library",
            "//vendor/github.com/docker/libnetwork/resolvconf:go_default_library",
            "//vendor/github.com/vishvananda/netlink:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = ["utils_test.go"],
    importpath = "github.com/docker/libnetwork/netutils",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
