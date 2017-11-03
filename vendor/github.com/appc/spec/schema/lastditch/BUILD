load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "image.go",
        "labels.go",
        "pod.go",
    ],
    importpath = "github.com/appc/spec/schema/lastditch",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "common_test.go",
        "image_test.go",
        "pod_test.go",
    ],
    importpath = "github.com/appc/spec/schema/lastditch",
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
