load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["etchosts.go"],
    importpath = "github.com/docker/libnetwork/etchosts",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["etchosts_test.go"],
    importpath = "github.com/docker/libnetwork/etchosts",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/libnetwork/testutils:go_default_library"],
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
