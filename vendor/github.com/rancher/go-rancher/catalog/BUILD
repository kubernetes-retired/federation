load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "common.go",
        "generated_api_version.go",
        "generated_catalog.go",
        "generated_client.go",
        "generated_error.go",
        "generated_question.go",
        "generated_template.go",
        "generated_template_version.go",
        "schemas.go",
        "types.go",
    ],
    importpath = "github.com/rancher/go-rancher/catalog",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gorilla/websocket:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
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
