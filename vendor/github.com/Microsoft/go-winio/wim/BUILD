load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "decompress.go",
        "wim.go",
    ],
    importpath = "github.com/Microsoft/go-winio/wim",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/Microsoft/go-winio/wim/lzx:go_default_library"],
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
        "//vendor/github.com/Microsoft/go-winio/wim/lzx:all-srcs",
        "//vendor/github.com/Microsoft/go-winio/wim/validate:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
