load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "error.go",
        "parser.go",
    ],
    importpath = "github.com/hashicorp/hcl/hcl/parser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/hcl/hcl/ast:go_default_library",
        "//vendor/github.com/hashicorp/hcl/hcl/scanner:go_default_library",
        "//vendor/github.com/hashicorp/hcl/hcl/token:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "error_test.go",
        "parser_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/hashicorp/hcl/hcl/parser",
    deps = [
        "//vendor/github.com/hashicorp/hcl/hcl/ast:go_default_library",
        "//vendor/github.com/hashicorp/hcl/hcl/token:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
