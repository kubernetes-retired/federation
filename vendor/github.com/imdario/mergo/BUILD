load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "map.go",
        "merge.go",
        "mergo.go",
    ],
    importpath = "github.com/imdario/mergo",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["mergo_test.go"],
    importpath = "github.com/imdario/mergo",
    library = ":go_default_library",
    deps = ["//vendor/gopkg.in/yaml.v1:go_default_library"],
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
