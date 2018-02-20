load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "decoder.go",
        "path.go",
        "pathaction.go",
    ],
    importpath = "github.com/exponent-io/jsonpath",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "decoder_test.go",
        "example_test.go",
        "path_test.go",
        "pathaction_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/exponent-io/jsonpath",
    deps = [
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
