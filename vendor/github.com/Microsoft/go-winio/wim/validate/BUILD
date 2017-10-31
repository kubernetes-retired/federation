load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["validate.go"],
    importpath = "github.com/Microsoft/go-winio/wim/validate",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/Microsoft/go-winio/wim:go_default_library"],
)

go_binary(
    name = "validate",
    importpath = "github.com/Microsoft/go-winio/wim/validate",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
