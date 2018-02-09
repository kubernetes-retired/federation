load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "error.go",
        "utils.go",
        "zfs.go",
        "zpool.go",
    ],
    importpath = "github.com/mistifyio/go-zfs",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pborman/uuid:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["error_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/mistifyio/go-zfs",
)

go_test(
    name = "go_default_xtest",
    srcs = ["zfs_test.go"],
    importpath = "github.com/mistifyio/go-zfs_test",
    deps = [":go_default_library"],
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
