load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["config.go"],
    importpath = "github.com/docker/libnetwork/config",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/go-connections/tlsconfig:go_default_library",
        "//vendor/github.com/docker/libkv/store:go_default_library",
        "//vendor/github.com/docker/libnetwork/cluster:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/osl:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["config_test.go"],
    importpath = "github.com/docker/libnetwork/config",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
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
