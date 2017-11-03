load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["enter_kvm.go"],
    importpath = "github.com/coreos/rkt/stage1/enter_kvm",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common/ssh:go_default_library",
    ],
)

go_binary(
    name = "enter_kvm",
    importpath = "github.com/coreos/rkt/stage1/enter_kvm",
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
