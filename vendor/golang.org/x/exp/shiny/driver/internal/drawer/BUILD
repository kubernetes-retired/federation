load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["drawer.go"],
    importpath = "golang.org/x/exp/shiny/driver/internal/drawer",
    visibility = ["//vendor/golang.org/x/exp/shiny/driver:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/exp/shiny/screen:go_default_library",
        "//vendor/golang.org/x/image/math/f64:go_default_library",
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
