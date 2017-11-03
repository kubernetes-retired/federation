load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "imports.go",
        "map.go",
        "methodsetcache.go",
        "ui.go",
    ],
    importpath = "golang.org/x/tools/go/types/typeutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "example_test.go",
        "imports_test.go",
        "map_test.go",
        "ui_test.go",
    ],
    importpath = "golang.org/x/tools/go/types/typeutil_test",
    deps = [":go_default_library"],
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
