load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ctl.go",
        "help.go",
    ],
    importpath = "github.com/coreos/etcd/etcdctl/ctlv3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv3/command:go_default_library",
        "//vendor/github.com/coreos/etcd/version:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv3/command:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
