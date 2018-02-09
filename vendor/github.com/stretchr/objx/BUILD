load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "accessors.go",
        "constants.go",
        "conversions.go",
        "doc.go",
        "map.go",
        "mutations.go",
        "security.go",
        "tests.go",
        "type_specific_codegen.go",
        "value.go",
    ],
    importpath = "github.com/stretchr/objx",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "accessors_test.go",
        "conversions_test.go",
        "fixture_test.go",
        "map_for_test.go",
        "map_test.go",
        "mutations_test.go",
        "security_test.go",
        "simple_example_test.go",
        "tests_test.go",
        "type_specific_codegen_test.go",
        "value_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/stretchr/objx",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
