load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["cors.go"],
    importpath = "github.com/coreos/etcd/pkg/cors",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["cors_test.go"],
    importpath = "github.com/coreos/etcd/pkg/cors",
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
