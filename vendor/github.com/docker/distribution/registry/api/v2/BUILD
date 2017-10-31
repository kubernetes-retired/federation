load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "descriptors.go",
        "doc.go",
        "errors.go",
        "headerparser.go",
        "routes.go",
        "urls.go",
    ],
    importpath = "github.com/docker/distribution/registry/api/v2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/distribution/registry/api/errcode:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "headerparser_test.go",
        "routes_test.go",
        "urls_test.go",
    ],
    importpath = "github.com/docker/distribution/registry/api/v2",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
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
