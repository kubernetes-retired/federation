load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fixtures.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/openstack/networking/v2/subnets/testing",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/subnets:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "requests_test.go",
        "results_test.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/openstack/networking/v2/subnets/testing",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/gophercloud/gophercloud:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/common:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/openstack/networking/v2/subnets:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/pagination:go_default_library",
        "//vendor/github.com/gophercloud/gophercloud/testhelper:go_default_library",
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
