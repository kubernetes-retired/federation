load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bypass.go",
        "common.go",
        "config.go",
        "doc.go",
        "dump.go",
        "format.go",
        "spew.go",
    ],
    importpath = "github.com/davecgh/go-spew/spew",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "internal_test.go",
        "internalunsafe_test.go",
    ],
    importpath = "github.com/davecgh/go-spew/spew",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "common_test.go",
        "dump_test.go",
        "dumpnocgo_test.go",
        "example_test.go",
        "format_test.go",
        "spew_test.go",
    ],
    importpath = "github.com/davecgh/go-spew/spew_test",
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
