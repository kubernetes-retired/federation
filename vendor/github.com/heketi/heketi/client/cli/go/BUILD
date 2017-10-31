load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/heketi/heketi/client/cli/go",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/heketi/heketi/client/cli/go/cmds:go_default_library"],
)

go_binary(
    name = "go",
    importpath = "github.com/heketi/heketi/client/cli/go",
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
        "//vendor/github.com/heketi/heketi/client/cli/go/cmds:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
