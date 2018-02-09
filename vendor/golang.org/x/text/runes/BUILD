load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cond.go",
        "runes.go",
    ],
    importpath = "golang.org/x/text/runes",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/text/transform:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cond_test.go",
        "runes_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "golang.org/x/text/runes",
    deps = [
        "//vendor/golang.org/x/text/cases:go_default_library",
        "//vendor/golang.org/x/text/internal/testtext:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "golang.org/x/text/runes_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
        "//vendor/golang.org/x/text/width:go_default_library",
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
