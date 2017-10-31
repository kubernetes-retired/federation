load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "ebnflint.go",
    ],
    importpath = "golang.org/x/exp/ebnflint",
    visibility = ["//visibility:private"],
    deps = ["//vendor/golang.org/x/exp/ebnf:go_default_library"],
)

go_binary(
    name = "ebnflint",
    importpath = "golang.org/x/exp/ebnflint",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["ebnflint_test.go"],
    importpath = "golang.org/x/exp/ebnflint",
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
