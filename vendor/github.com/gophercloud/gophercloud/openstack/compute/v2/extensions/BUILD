load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "delegate.go",
        "doc.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/common/extensions:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/pagination:go_default_library",
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
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/attachinterfaces:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/availabilityzones:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/defsecrules:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/diskconfig:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/evacuate:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/limits:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/lockunlock:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/migrate:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/networks:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/pauseunpause:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/resetstate:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/schedulerhints:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/suspendresume:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/tenantnetworks:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/testing:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
