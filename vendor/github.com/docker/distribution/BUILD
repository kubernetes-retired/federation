load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "blobs.go",
        "doc.go",
        "errors.go",
        "manifests.go",
        "registry.go",
        "tags.go",
    ],
    importpath = "github.com/docker/distribution",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/distribution/cmd/digest:all-srcs",
        "//vendor/github.com/docker/distribution/cmd/registry:all-srcs",
        "//vendor/github.com/docker/distribution/cmd/registry-api-descriptor-template:all-srcs",
        "//vendor/github.com/docker/distribution/configuration:all-srcs",
        "//vendor/github.com/docker/distribution/context:all-srcs",
        "//vendor/github.com/docker/distribution/contrib/token-server:all-srcs",
        "//vendor/github.com/docker/distribution/digestset:all-srcs",
        "//vendor/github.com/docker/distribution/health:all-srcs",
        "//vendor/github.com/docker/distribution/manifest:all-srcs",
        "//vendor/github.com/docker/distribution/notifications:all-srcs",
        "//vendor/github.com/docker/distribution/reference:all-srcs",
        "//vendor/github.com/docker/distribution/registry:all-srcs",
        "//vendor/github.com/docker/distribution/testutil:all-srcs",
        "//vendor/github.com/docker/distribution/uuid:all-srcs",
        "//vendor/github.com/docker/distribution/version:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
