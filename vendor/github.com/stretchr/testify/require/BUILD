load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "forward_requirements.go",
        "require.go",
        "require_forward.go",
        "requirements.go",
    ],
    importpath = "github.com/stretchr/testify/require",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "forward_requirements_test.go",
        "requirements_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/stretchr/testify/require",
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
