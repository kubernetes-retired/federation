load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["netexec.go"],
    importpath = "k8s.io/kubernetes/test/images/netexec",
    visibility = ["//visibility:private"],
    deps = ["//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library"],
)

go_binary(
    name = "netexec",
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/test/images/netexec",
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
