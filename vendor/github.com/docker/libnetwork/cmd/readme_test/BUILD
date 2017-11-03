load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["readme.go"],
    importpath = "github.com/docker/libnetwork/cmd/readme_test",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork:go_default_library",
        "//vendor/github.com/docker/libnetwork/config:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
    ],
)

go_binary(
    name = "readme_test",
    importpath = "github.com/docker/libnetwork/cmd/readme_test",
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
