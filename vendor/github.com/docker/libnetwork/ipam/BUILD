load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "allocator.go",
        "store.go",
        "structures.go",
        "utils.go",
    ],
    importpath = "github.com/docker/libnetwork/ipam",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/bitseq:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["allocator_test.go"],
    importpath = "github.com/docker/libnetwork/ipam",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libkv/store:go_default_library",
        "//vendor/github.com/docker/libkv/store/boltdb:go_default_library",
        "//vendor/github.com/docker/libnetwork/bitseq:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
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
