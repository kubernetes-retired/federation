load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bridge.go",
        "dispatch.go",
    ],
    importpath = "github.com/coreos/etcd/tools/local-tester/bridge",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "bridge",
    importpath = "github.com/coreos/etcd/tools/local-tester/bridge",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
