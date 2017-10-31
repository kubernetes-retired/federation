load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/libopenstorage/openstorage/pkg/flexvolume/cmd/flexvolume",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/pkg/flexvolume:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/go.pedge.io/env:go_default_library",
        "//vendor/go.pedge.io/lion/env:go_default_library",
        "//vendor/go.pedge.io/pkg/cobra:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)

go_binary(
    name = "flexvolume",
    importpath = "github.com/libopenstorage/openstorage/pkg/flexvolume/cmd/flexvolume",
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
