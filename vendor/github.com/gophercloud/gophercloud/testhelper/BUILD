load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "convenience.go",
        "doc.go",
        "http_responses.go",
    ],
    importpath = "github.com/gophercloud/gophercloud/testhelper",
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
        "//vendor/github.com/gophercloud/gophercloud/testhelper/client:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/testhelper/fixture:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
