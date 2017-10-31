load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/coreos/etcd/etcdctl",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv2:go_default_library",
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv3:go_default_library",
    ],
)

go_binary(
    name = "etcdctl",
    importpath = "github.com/coreos/etcd/etcdctl",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv2:all-srcs",
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv3:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
