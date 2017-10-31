load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["common.go"],
    importpath = "github.com/gophercloud/gophercloud/acceptance/openstack",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gophercloud/gophercloud/openstack/common/extensions:go_default_library"],
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
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/blockstorage/extensions:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/blockstorage/noauth:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/blockstorage/v1:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/blockstorage/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/compute/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/db/v1:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/dns/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/identity/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/identity/v3:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/imageservice/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/objectstorage/v1:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/orchestration/v1:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/sharedfilesystems/v2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
