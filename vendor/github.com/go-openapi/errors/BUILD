load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "auth.go",
        "headers.go",
        "middleware.go",
        "parsing.go",
        "schema.go",
    ],
    importpath = "github.com/go-openapi/errors",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "api_test.go",
        "auth_test.go",
        "middleware_test.go",
        "parsing_test.go",
        "schema_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/go-openapi/errors",
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
