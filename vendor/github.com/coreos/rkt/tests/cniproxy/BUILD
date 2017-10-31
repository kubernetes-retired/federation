load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/coreos/rkt/tests/cniproxy",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/containernetworking/cni/pkg/invoke:go_default_library"],
)

go_binary(
    name = "cniproxy",
    importpath = "github.com/coreos/rkt/tests/cniproxy",
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
