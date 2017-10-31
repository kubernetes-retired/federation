load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "arch.go",
        "attach.go",
        "common.go",
        "log.go",
        "registration.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "app_add.go",
            "app_rm.go",
            "app_start.go",
            "app_stop.go",
            "enter.go",
            "gc.go",
            "manifest.go",
            "run.go",
            "stop.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/stage0",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/common/apps:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common/types:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
            "//vendor/github.com/coreos/rkt/common/networking:go_default_library",
            "//vendor/github.com/coreos/rkt/common/overlay:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/aci:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/label:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/pod:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/tpm:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
            "//vendor/github.com/coreos/rkt/store/imagestore:go_default_library",
            "//vendor/github.com/coreos/rkt/store/treestore:go_default_library",
            "//vendor/github.com/coreos/rkt/version:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
