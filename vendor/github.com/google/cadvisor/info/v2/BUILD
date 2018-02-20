load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "container.go",
        "conversion.go",
        "machine.go",
    ],
    importpath = "github.com/google/cadvisor/info/v2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["conversion_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/google/cadvisor/info/v2",
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
    ],
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
