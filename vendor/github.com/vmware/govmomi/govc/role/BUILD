load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "create.go",
        "ls.go",
        "remove.go",
        "update.go",
        "usage.go",
    ],
    importpath = "github.com/vmware/govmomi/govc/role",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/govc/cli:go_default_library",
        "//vendor/github.com/vmware/govmomi/govc/permissions:go_default_library",
        "//vendor/github.com/vmware/govmomi/object:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
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
