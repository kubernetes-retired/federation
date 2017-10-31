load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "lbaas.go",
        "pkg.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/lbaas",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/tools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/members:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/monitors:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/pools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas/vips:go_default_library",
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
