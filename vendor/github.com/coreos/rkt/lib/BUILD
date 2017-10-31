load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "app.go",
        "doc.go",
        "pod.go",
        "types.go",
    ],
    importpath = "github.com/coreos/rkt/lib",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/networking/netinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/pod:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
