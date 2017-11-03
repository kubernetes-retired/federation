load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["flex.go"],
    importpath = "golang.org/x/exp/shiny/widget/flex",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/widget/node:go_default_library",
        "//vendor/golang.org/x/exp/shiny/widget/theme:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["flex_test.go"],
    importpath = "golang.org/x/exp/shiny/widget/flex",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/exp/shiny/unit:go_default_library",
        "//vendor/golang.org/x/exp/shiny/widget:go_default_library",
        "//vendor/golang.org/x/exp/shiny/widget/node:go_default_library",
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
