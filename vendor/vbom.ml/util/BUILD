load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "regexp.go",
    ],
    importpath = "vbom.ml/util",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["regexp_test.go"],
    importpath = "vbom.ml/util",
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
    srcs = [
        ":package-srcs",
        "//vendor/vbom.ml/util/cmd/short-regexp:all-srcs",
        "//vendor/vbom.ml/util/rope:all-srcs",
        "//vendor/vbom.ml/util/sortorder:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
