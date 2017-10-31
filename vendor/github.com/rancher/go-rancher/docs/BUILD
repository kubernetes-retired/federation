load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "api_description_generator.go",
        "api_doc_generator.go",
        "api_general_functions.go",
        "main.go",
    ],
    importpath = "github.com/rancher/go-rancher/docs",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/rancher/go-rancher/client:go_default_library",
        "//vendor/gopkg.in/yaml.v2:go_default_library",
    ],
)

go_binary(
    name = "docs",
    importpath = "github.com/rancher/go-rancher/docs",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
