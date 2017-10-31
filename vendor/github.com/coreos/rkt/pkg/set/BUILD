load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["string.go"],
    importpath = "github.com/coreos/rkt/pkg/set",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["string_test.go"],
    importpath = "github.com/coreos/rkt/pkg/set",
    library = ":go_default_library",
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
