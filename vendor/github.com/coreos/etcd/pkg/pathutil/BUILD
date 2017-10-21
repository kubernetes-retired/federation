load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["path.go"],
    importpath = "github.com/coreos/etcd/pkg/pathutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["path_test.go"],
    importpath = "github.com/coreos/etcd/pkg/pathutil",
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
