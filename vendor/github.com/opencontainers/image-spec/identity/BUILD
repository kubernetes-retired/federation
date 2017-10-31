load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "chainid.go",
        "helpers.go",
    ],
    importpath = "github.com/opencontainers/image-spec/identity",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/opencontainers/go-digest:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["chainid_test.go"],
    importpath = "github.com/opencontainers/image-spec/identity",
    library = ":go_default_library",
    deps = ["//vendor/github.com/opencontainers/go-digest:go_default_library"],
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
