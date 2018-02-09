load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["coredns.go"],
    importpath = "github.com/coredns/coredns",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/coredns/coredns/core/plugin:go_default_library",
        "//vendor/github.com/coredns/coredns/coremain:go_default_library",
    ],
)

go_binary(
    name = "coredns",
    embed = [":go_default_library"],
    importpath = "github.com/coredns/coredns",
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
        "//vendor/github.com/coredns/coredns/plugin/etcd/msg:all-srcs",
        "//vendor/github.com/coredns/coredns/plugin/pkg/dnsutil:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
