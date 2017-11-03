load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "graph.go",
        "memory_graph.go",
        "statement.go",
    ],
    importpath = "github.com/docker/libtrust/trustgraph",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/libtrust:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "memory_graph_test.go",
        "statement_test.go",
    ],
    importpath = "github.com/docker/libtrust/trustgraph",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/libtrust:go_default_library",
        "//vendor/github.com/docker/libtrust/testutil:go_default_library",
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
