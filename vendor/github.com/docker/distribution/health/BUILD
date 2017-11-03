load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "health.go",
    ],
    importpath = "github.com/docker/distribution/health",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/registry/api/errcode:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["health_test.go"],
    importpath = "github.com/docker/distribution/health",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/distribution/health/api:all-srcs",
        "//vendor/github.com/docker/distribution/health/checks:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
