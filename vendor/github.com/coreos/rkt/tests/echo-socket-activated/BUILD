load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["echo-socket-activated.go"],
    importpath = "github.com/coreos/rkt/tests/echo-socket-activated",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/coreos/rkt/pkg/sys:go_default_library"],
)

go_binary(
    name = "echo-socket-activated",
    importpath = "github.com/coreos/rkt/tests/echo-socket-activated",
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
