load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cobra.go",
        "error.go",
        "required.go",
    ],
    importpath = "github.com/docker/docker/cli",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/docker/cli/config:all-srcs",
        "//vendor/github.com/docker/docker/cli/debug:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
