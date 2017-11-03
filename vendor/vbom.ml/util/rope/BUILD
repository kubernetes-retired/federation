load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "concat.go",
        "debug.go",
        "leaf.go",
        "node.go",
        "reader.go",
        "rebalance.go",
        "rope.go",
    ],
    importpath = "vbom.ml/util/rope",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "concat_test.go",
        "leaf_test.go",
        "node_test.go",
        "reader_test.go",
        "rebalance_test.go",
        "rope_test.go",
    ],
    importpath = "vbom.ml/util/rope",
    library = ":go_default_library",
    deps = ["//vendor/github.com/bruth/assert:go_default_library"],
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
