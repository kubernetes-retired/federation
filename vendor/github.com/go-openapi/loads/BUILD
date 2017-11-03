load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["spec.go"],
    importpath = "github.com/go-openapi/loads",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/go-openapi/analysis:go_default_library",
        "//vendor/github.com/go-openapi/spec:go_default_library",
        "//vendor/github.com/go-openapi/swag:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "json_test.go",
        "spec_test.go",
    ],
    importpath = "github.com/go-openapi/loads",
    library = ":go_default_library",
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
        "//vendor/github.com/go-openapi/loads/fmts:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
