load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "formula.go",
        "table.go",
    ],
    importpath = "github.com/chai2010/gettext-go/gettext/plural",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["formula_test.go"],
    importpath = "github.com/chai2010/gettext-go/gettext/plural",
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
