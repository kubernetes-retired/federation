load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "ratelimit.go",
        "reader.go",
    ],
    importpath = "github.com/juju/ratelimit",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["ratelimit_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/juju/ratelimit",
    deps = ["//vendor/gopkg.in/check.v1:go_default_library"],
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
