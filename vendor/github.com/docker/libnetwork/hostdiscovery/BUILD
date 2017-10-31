load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "hostdiscovery.go",
        "hostdiscovery_api.go",
    ],
    importpath = "github.com/docker/libnetwork/hostdiscovery",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/deckarep/golang-set:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery/kv:go_default_library",
        "//vendor/github.com/docker/libkv/store/consul:go_default_library",
        "//vendor/github.com/docker/libkv/store/etcd:go_default_library",
        "//vendor/github.com/docker/libkv/store/zookeeper:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["hostdiscovery_test.go"],
    importpath = "github.com/docker/libnetwork/hostdiscovery",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/deckarep/golang-set:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
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
