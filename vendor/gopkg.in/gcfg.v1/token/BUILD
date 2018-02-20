load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "position.go",
        "serialize.go",
        "token.go",
    ],
    importpath = "gopkg.in/gcfg.v1/token",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "position_test.go",
        "serialize_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "gopkg.in/gcfg.v1/token",
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
