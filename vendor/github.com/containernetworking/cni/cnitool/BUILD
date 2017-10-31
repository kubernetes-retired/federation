load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cni.go"],
    importpath = "github.com/containernetworking/cni/cnitool",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/containernetworking/cni/libcni:go_default_library"],
)

go_binary(
    name = "cnitool",
    importpath = "github.com/containernetworking/cni/cnitool",
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
