load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "extensions.go",
        "pkg.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/tools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/external:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/networks:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/ports:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/fwaas:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/layer3:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/lbaas:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/lbaas_v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/portsbinding:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
