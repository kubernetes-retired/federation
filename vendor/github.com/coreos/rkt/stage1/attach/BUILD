load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["attach.go"],
    importpath = "github.com/coreos/rkt/stage1/attach",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/init/common:go_default_library",
    ],
)

go_binary(
    name = "attach",
    importpath = "github.com/coreos/rkt/stage1/attach",
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
