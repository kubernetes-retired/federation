load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "buffer.go",
        "say_matcher.go",
    ],
    importpath = "github.com/onsi/gomega/gbytes",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/onsi/gomega/format:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "buffer_test.go",
        "gbuffer_suite_test.go",
        "say_matcher_test.go",
    ],
    importpath = "github.com/onsi/gomega/gbytes_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
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
