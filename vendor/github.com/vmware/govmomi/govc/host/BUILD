load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "add.go",
        "disconnect.go",
        "info.go",
        "reconnect.go",
        "remove.go",
    ],
    importpath = "github.com/vmware/govmomi/govc/host",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/govc/cli:go_default_library",
        "//vendor/github.com/vmware/govmomi/govc/flags:go_default_library",
        "//vendor/github.com/vmware/govmomi/object:go_default_library",
        "//vendor/github.com/vmware/govmomi/property:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/methods:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/mo:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/account:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/autostart:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/cert:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/date:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/esxcli:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/firewall:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/maintenance:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/option:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/portgroup:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/service:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/storage:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/vnic:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/host/vswitch:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
