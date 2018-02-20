load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "assertion_format.go",
        "assertion_forward.go",
        "assertions.go",
        "doc.go",
        "errors.go",
        "forward_assertions.go",
        "http_assertions.go",
    ],
    importpath = "github.com/stretchr/testify/assert",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/pmezard/go-difflib/difflib:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "assertions_test.go",
        "forward_assertions_test.go",
        "http_assertions_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/stretchr/testify/assert",
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
