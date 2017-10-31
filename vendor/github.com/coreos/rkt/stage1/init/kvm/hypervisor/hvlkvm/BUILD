load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["lkvm_driver.go"],
    importpath = "github.com/coreos/rkt/stage1/init/kvm/hypervisor/hvlkvm",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/stage1/init/kvm:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor:go_default_library",
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
