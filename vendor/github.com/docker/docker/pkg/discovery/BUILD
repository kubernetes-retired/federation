load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "backends.go",
        "discovery.go",
        "entry.go",
        "generator.go",
    ],
    importpath = "github.com/docker/docker/pkg/discovery",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/sirupsen/logrus:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "discovery_test.go",
        "generator_test.go",
    ],
    importpath = "github.com/docker/docker/pkg/discovery",
    library = ":go_default_library",
    deps = ["//vendor/github.com/go-check/check:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/file:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/kv:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/memory:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/nodes:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
