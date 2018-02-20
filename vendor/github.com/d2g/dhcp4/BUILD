load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "constants.go",
        "helpers.go",
        "option.go",
        "packet.go",
    ],
    importpath = "github.com/d2g/dhcp4",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "helpers_test.go",
        "packet_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/d2g/dhcp4",
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
