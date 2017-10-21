load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "base.go",
        "cldr.go",
        "collate.go",
        "decode.go",
        "resolve.go",
        "slice.go",
        "xml.go",
    ],
    importpath = "golang.org/x/text/unicode/cldr",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cldr_test.go",
        "collate_test.go",
        "data_test.go",
        "resolve_test.go",
        "slice_test.go",
    ],
    importpath = "golang.org/x/text/unicode/cldr",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["examples_test.go"],
    importpath = "golang.org/x/text/unicode/cldr_test",
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
