load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "chroot.go",
        "tar.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/tar",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/pkg/device:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/multicall:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["tar_test.go"],
    importpath = "github.com/coreos/rkt/pkg/tar",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/multicall:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
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
