load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["lifecycler.go"],
    importpath = "golang.org/x/exp/shiny/driver/internal/lifecycler",
    visibility = ["//vendor/golang.org/x/exp/shiny/driver:__subpackages__"],
    deps = ["//vendor/golang.org/x/mobile/event/lifecycle:go_default_library"],
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
