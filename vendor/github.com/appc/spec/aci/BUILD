load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "build.go",
        "doc.go",
        "file.go",
        "layout.go",
        "writer.go",
    ],
    importpath = "github.com/appc/spec/aci",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/pkg/tarheader:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "file_test.go",
        "layout_test.go",
    ],
    importpath = "github.com/appc/spec/aci",
    library = ":go_default_library",
    deps = ["//vendor/github.com/appc/spec/schema:go_default_library"],
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
