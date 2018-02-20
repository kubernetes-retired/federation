load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "capture.go",
        "frame.go",
        "stacktrace.go",
    ],
    importpath = "github.com/opencontainers/runc/libcontainer/stacktrace",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "capture_test.go",
        "frame_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/opencontainers/runc/libcontainer/stacktrace",
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
