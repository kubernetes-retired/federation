load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "enclosing.go",
        "imports.go",
        "util.go",
    ],
    importpath = "golang.org/x/tools/go/ast/astutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["imports_test.go"],
    importpath = "golang.org/x/tools/go/ast/astutil",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["enclosing_test.go"],
    importpath = "golang.org/x/tools/go/ast/astutil_test",
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
