load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/coreos/rkt/stage1_fly/run",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/flag:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fs:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/mountinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/common/types:go_default_library",
        "//vendor/github.com/coreos/rkt/stage1/init/common:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_binary(
    name = "run",
    importpath = "github.com/coreos/rkt/stage1_fly/run",
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
