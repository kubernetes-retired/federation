load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "mount.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "mount_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/pkg/fs",
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
