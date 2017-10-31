load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cni_args.go",
        "raw_exec.go",
        "version_decoder.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/invoke/fakes",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/containernetworking/cni/pkg/version:go_default_library"],
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
