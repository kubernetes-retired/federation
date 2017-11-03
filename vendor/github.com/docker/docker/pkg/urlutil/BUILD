load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["urlutil.go"],
    importpath = "github.com/docker/docker/pkg/urlutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["urlutil_test.go"],
    importpath = "github.com/docker/docker/pkg/urlutil",
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
