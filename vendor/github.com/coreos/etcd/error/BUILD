load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["error.go"],
    importpath = "github.com/coreos/etcd/error",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["error_test.go"],
    importpath = "github.com/coreos/etcd/error",
    library = ":go_default_library",
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
