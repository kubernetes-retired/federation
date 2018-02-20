load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["example.go"],
    importpath = "github.com/coreos/go-semver",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/coreos/go-semver/semver:go_default_library"],
)

go_binary(
    name = "go-semver",
    embed = [":go_default_library"],
    importpath = "github.com/coreos/go-semver",
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
        "//vendor/github.com/coreos/go-semver/semver:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
