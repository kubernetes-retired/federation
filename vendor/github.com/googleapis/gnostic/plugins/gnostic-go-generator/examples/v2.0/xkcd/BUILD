load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/xkcd",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/xkcd/xkcd:go_default_library"],
)

go_binary(
    name = "xkcd",
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/xkcd",
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
