load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "namespace_unsupported.go",
        "sandbox.go",
        "sandbox_unsupported.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "interface_linux.go",
            "namespace_linux.go",
            "neigh_linux.go",
            "options_linux.go",
            "route_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "interface_windows.go",
            "namespace_windows.go",
            "neigh_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/osl",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
            "//vendor/github.com/docker/libnetwork/ns:go_default_library",
            "//vendor/github.com/sirupsen/logrus:go_default_library",
            "//vendor/github.com/vishvananda/netlink:go_default_library",
            "//vendor/github.com/vishvananda/netns:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = [
        "sandbox_test.go",
        "sandbox_unsupported_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "sandbox_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/osl",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork/ns:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/libnetwork/types:go_default_library",
            "//vendor/github.com/vishvananda/netlink:go_default_library",
            "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
            "//vendor/github.com/vishvananda/netns:go_default_library",
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
