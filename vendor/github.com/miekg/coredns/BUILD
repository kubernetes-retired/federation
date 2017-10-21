load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["coredns.go"],
    importpath = "github.com/miekg/coredns",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/miekg/coredns/coremain:go_default_library"],
)

go_binary(
    name = "coredns",
    importpath = "github.com/miekg/coredns",
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
        "//vendor/github.com/miekg/coredns/middleware:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
