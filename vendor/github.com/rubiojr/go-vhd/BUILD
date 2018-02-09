load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cli.go",
        "version.go",
    ],
    importpath = "github.com/rubiojr/go-vhd",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/codegangsta/cli:go_default_library",
        "//vendor/github.com/dustin/go-humanize:go_default_library",
        "//vendor/github.com/rubiojr/go-vhd/vhd:go_default_library",
    ],
)

go_binary(
    name = "go-vhd",
    embed = [":go_default_library"],
    importpath = "github.com/rubiojr/go-vhd",
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
        "//vendor/github.com/rubiojr/go-vhd/vhd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
