load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "flow.go",
        "image.go",
        "label.go",
        "text.go",
        "uniform.go",
        "widget.go",
    ],
    importpath = "golang.org/x/exp/shiny/widget",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/gesture:go_default_library",
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
        "//vendor/golang.org/x/exp/shiny/text:go_default_library",
        "//vendor/golang.org/x/exp/shiny/unit:go_default_library",
        "//vendor/golang.org/x/exp/shiny/widget/node:go_default_library",
        "//vendor/golang.org/x/exp/shiny/widget/theme:go_default_library",
        "//vendor/golang.org/x/image/font:go_default_library",
        "//vendor/golang.org/x/image/math/fixed:go_default_library",
        "//vendor/golang.org/x/mobile/event/lifecycle:go_default_library",
        "//vendor/golang.org/x/mobile/event/mouse:go_default_library",
        "//vendor/golang.org/x/mobile/event/paint:go_default_library",
        "//vendor/golang.org/x/mobile/event/size:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/exp/shiny/widget/flex:all-srcs",
        "//vendor/golang.org/x/exp/shiny/widget/node:all-srcs",
        "//vendor/golang.org/x/exp/shiny/widget/theme:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
