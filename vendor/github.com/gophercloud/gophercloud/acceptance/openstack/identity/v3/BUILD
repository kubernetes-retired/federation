load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "identity.go",
        "pkg.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/acceptance/openstack/identity/v3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/tools:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/domains:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/groups:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/projects:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/identity/v3/users:go_default_library",
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
