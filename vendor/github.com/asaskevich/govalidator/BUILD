load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "arrays.go",
        "converter.go",
        "error.go",
        "numerics.go",
        "patterns.go",
        "types.go",
        "utils.go",
        "validator.go",
    ],
    importpath = "github.com/asaskevich/govalidator",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "arrays_test.go",
        "converter_test.go",
        "error_test.go",
        "numerics_test.go",
        "utils_test.go",
        "validator_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/asaskevich/govalidator",
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
