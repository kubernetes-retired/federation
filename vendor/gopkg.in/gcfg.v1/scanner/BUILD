load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "errors.go",
        "scanner.go",
    ],
    importpath = "gopkg.in/gcfg.v1/scanner",
    visibility = ["//visibility:public"],
    deps = ["//vendor/gopkg.in/gcfg.v1/token:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["scanner_test.go"],
    embed = [":go_default_library"],
    importpath = "gopkg.in/gcfg.v1/scanner",
    deps = ["//vendor/gopkg.in/gcfg.v1/token:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "gopkg.in/gcfg.v1/scanner_test",
    deps = [
        ":go_default_library",
        "//vendor/gopkg.in/gcfg.v1/token:go_default_library",
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
