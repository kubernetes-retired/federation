load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "ct.go",
        "ct_ansi.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "ct_win.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/daviddengcn/go-colortext",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "ct_ansi_test.go",
        "ct_test.go",
    ],
    importpath = "github.com/daviddengcn/go-colortext",
    library = ":go_default_library",
    deps = ["//vendor/github.com/golangplus/testing/assert:go_default_library"],
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
