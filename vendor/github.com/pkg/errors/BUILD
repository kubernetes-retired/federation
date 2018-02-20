load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "errors.go",
        "stack.go",
    ],
    importpath = "github.com/pkg/errors",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bench_test.go",
        "errors_test.go",
        "format_test.go",
        "stack_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/pkg/errors",
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "github.com/pkg/errors_test",
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
