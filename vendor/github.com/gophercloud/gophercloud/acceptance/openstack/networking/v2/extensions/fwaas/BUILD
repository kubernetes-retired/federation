load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fwaas.go",
        "pkg.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/fwaas",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/tools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/firewalls:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/policies:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/routerinsertion:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules:go_default_library",
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
