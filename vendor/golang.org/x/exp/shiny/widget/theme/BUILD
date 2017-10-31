load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["theme.go"],
    importpath = "golang.org/x/exp/shiny/widget/theme",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/unit:go_default_library",
        "//vendor/golang.org/x/image/font:go_default_library",
        "//vendor/golang.org/x/image/font/basicfont:go_default_library",
        "//vendor/golang.org/x/image/math/fixed:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["theme_test.go"],
    importpath = "golang.org/x/exp/shiny/widget/theme",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/exp/shiny/unit:go_default_library",
        "//vendor/golang.org/x/image/math/fixed:go_default_library",
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
