load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["types.go"],
    importpath = "github.com/json-iterator/go/output_tests/struct/ptr_int32",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["json_test.go"],
    importpath = "github.com/json-iterator/go/output_tests/struct/ptr_int32",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/google/gofuzz:go_default_library",
        "//vendor/github.com/json-iterator/go:go_default_library",
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
