load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "byte.go",
        "doc.go",
        "empty.go",
        "int.go",
        "int64.go",
        "string.go",
    ],
    importpath = "k8s.io/gengo/examples/set-gen/sets",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["set_test.go"],
    embed = [":go_default_library"],
    importpath = "k8s.io/gengo/examples/set-gen/sets",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/k8s.io/gengo/examples/set-gen/sets/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
