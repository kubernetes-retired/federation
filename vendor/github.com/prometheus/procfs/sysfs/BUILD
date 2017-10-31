load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fs.go",
    ],
    importpath = "github.com/prometheus/procfs/sysfs",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/prometheus/procfs/xfs:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["fs_test.go"],
    importpath = "github.com/prometheus/procfs/sysfs",
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
