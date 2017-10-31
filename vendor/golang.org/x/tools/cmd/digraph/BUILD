load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["digraph.go"],
    importpath = "golang.org/x/tools/cmd/digraph",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "digraph",
    importpath = "golang.org/x/tools/cmd/digraph",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["digraph_test.go"],
    importpath = "golang.org/x/tools/cmd/digraph",
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
