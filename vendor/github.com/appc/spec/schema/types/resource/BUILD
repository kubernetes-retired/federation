load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "amount.go",
        "math.go",
        "quantity.go",
        "scale_int.go",
        "suffix.go",
    ],
    importpath = "github.com/appc/spec/schema/types/resource",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/gopkg.in/inf.v0:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "amount_test.go",
        "math_test.go",
        "quantity_test.go",
        "scale_int_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/appc/spec/schema/types/resource",
    deps = [
        "//vendor/github.com/google/gofuzz:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/gopkg.in/inf.v0:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["quantity_example_test.go"],
    importpath = "github.com/appc/spec/schema/types/resource_test",
    deps = [":go_default_library"],
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
