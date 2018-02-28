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
