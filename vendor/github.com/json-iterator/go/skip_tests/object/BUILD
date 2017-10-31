load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["inputs.go"],
    importpath = "github.com/json-iterator/go/skip_tests/object",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["skip_test.go"],
    importpath = "github.com/json-iterator/go/skip_tests/object",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/json-iterator/go:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
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
