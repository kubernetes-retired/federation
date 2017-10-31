load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "eval.go",
        "jsonconfig.go",
    ],
    importpath = "go4.org/jsonconfig",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go4.org/errorutil:go_default_library",
        "//vendor/go4.org/wkfs:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["jsonconfig_test.go"],
    data = glob(["testdata/**"]),
    importpath = "go4.org/jsonconfig",
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
