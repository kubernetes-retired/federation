load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cache.go",
        "datastore.go",
        "mock_store.go",
    ],
    importpath = "github.com/docker/libnetwork/datastore",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libkv:go_default_library",
        "//vendor/github.com/docker/libkv/store:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["datastore_test.go"],
    importpath = "github.com/docker/libnetwork/datastore",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
