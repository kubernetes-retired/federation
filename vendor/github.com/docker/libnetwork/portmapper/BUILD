load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "mapper.go",
        "mock_proxy.go",
        "proxy.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "proxy_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/portmapper",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/iptables:go_default_library",
        "//vendor/github.com/docker/libnetwork/portallocator:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["mapper_test.go"],
    importpath = "github.com/docker/libnetwork/portmapper",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libnetwork/iptables:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
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
