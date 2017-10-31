load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "versioned.go",
    ],
    importpath = "github.com/docker/distribution/manifest",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/distribution/manifest/manifestlist:all-srcs",
        "//vendor/github.com/docker/distribution/manifest/schema1:all-srcs",
        "//vendor/github.com/docker/distribution/manifest/schema2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
