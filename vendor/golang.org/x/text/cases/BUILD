load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cases.go",
        "context.go",
        "fold.go",
        "info.go",
        "map.go",
        "tables.go",
        "trieval.go",
    ],
    cgo = True,
    importpath = "golang.org/x/text/cases",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/text/internal:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "context_test.go",
        "fold_test.go",
        "map_test.go",
        "tables_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "golang.org/x/text/cases",
    deps = [
        "//vendor/golang.org/x/text/internal/testtext:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
        "//vendor/golang.org/x/text/unicode/rangetable:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "golang.org/x/text/cases_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
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
