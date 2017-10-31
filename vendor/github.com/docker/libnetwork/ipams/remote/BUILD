load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["remote.go"],
    importpath = "github.com/docker/libnetwork/ipams/remote",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugins:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipams/remote/api:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["remote_test.go"],
    importpath = "github.com/docker/libnetwork/ipams/remote",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugins:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/libnetwork/ipams/remote/api:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
