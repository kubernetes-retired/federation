load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
