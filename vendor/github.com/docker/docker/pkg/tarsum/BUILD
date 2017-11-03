load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "builder_context.go",
        "fileinfosums.go",
        "tarsum.go",
        "versioning.go",
        "writercloser.go",
    ],
    importpath = "github.com/docker/docker/pkg/tarsum",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "builder_context_test.go",
        "fileinfosums_test.go",
        "tarsum_test.go",
        "versioning_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "github.com/docker/docker/pkg/tarsum",
    library = ":go_default_library",
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
