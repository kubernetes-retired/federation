load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["backup.go"],
    importpath = "github.com/coreos/rkt/pkg/backup",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["backup_test.go"],
    importpath = "github.com/coreos/rkt/pkg/backup",
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
