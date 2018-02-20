load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["units.go"],
    importpath = "github.com/libopenstorage/openstorage/pkg/units",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["units_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/libopenstorage/openstorage/pkg/units",
    deps = ["//vendor/github.com/stretchr/testify/require:go_default_library"],
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
