load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "helpers.go",
        "normalize.go",
        "reference.go",
        "regexp.go",
    ],
    importpath = "github.com/docker/distribution/reference",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/digestset:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "normalize_test.go",
        "reference_test.go",
        "regexp_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/docker/distribution/reference",
    deps = [
        "//vendor/github.com/docker/distribution/digestset:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
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
