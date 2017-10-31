load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "auth_options.go",
        "doc.go",
        "endpoint_search.go",
        "errors.go",
        "params.go",
        "provider_client.go",
        "results.go",
        "service_client.go",
        "util.go",
    ],
    importpath = "github.com/gophercloud/gophercloud",
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
        "//vendor/github.com/gophercloud/gophercloud/acceptance/clients:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/tools:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/internal:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/openstack:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/pagination:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/testhelper:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/testing:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
