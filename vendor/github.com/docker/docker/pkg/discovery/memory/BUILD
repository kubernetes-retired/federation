load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["memory.go"],
    importpath = "github.com/docker/docker/pkg/discovery/memory",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/discovery:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["memory_test.go"],
    importpath = "github.com/docker/docker/pkg/discovery/memory",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
        "//vendor/github.com/go-check/check:go_default_library",
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
