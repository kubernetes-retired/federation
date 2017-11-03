load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "benchcmp.go",
        "compare.go",
        "doc.go",
    ],
    importpath = "golang.org/x/tools/cmd/benchcmp",
    visibility = ["//visibility:private"],
    deps = ["//vendor/golang.org/x/tools/benchmark/parse:go_default_library"],
)

go_binary(
    name = "benchcmp",
    importpath = "golang.org/x/tools/cmd/benchcmp",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "benchcmp_test.go",
        "compare_test.go",
    ],
    importpath = "golang.org/x/tools/cmd/benchcmp",
    library = ":go_default_library",
    deps = ["//vendor/golang.org/x/tools/benchmark/parse:go_default_library"],
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
