load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "context.go",
        "handler.go",
        "html_writer.go",
        "json_writer.go",
        "url.go",
        "writer.go",
    ],
    importpath = "github.com/rancher/go-rancher/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/gorilla/context:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/rancher/go-rancher/client:go_default_library",
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
