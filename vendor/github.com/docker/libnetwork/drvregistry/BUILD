load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["drvregistry.go"],
    importpath = "github.com/docker/libnetwork/drvregistry",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["drvregistry_test.go"],
    importpath = "github.com/docker/libnetwork/drvregistry",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipams/builtin:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipams/null:go_default_library",
        "//vendor/github.com/docker/libnetwork/ipams/remote:go_default_library",
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
