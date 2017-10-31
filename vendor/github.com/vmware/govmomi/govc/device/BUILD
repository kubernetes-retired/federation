load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "boot.go",
        "connect.go",
        "disconnect.go",
        "info.go",
        "ls.go",
        "remove.go",
    ],
    importpath = "github.com/vmware/govmomi/govc/device",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/govc/cli:go_default_library",
        "//vendor/github.com/vmware/govmomi/govc/flags:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/vmware/govmomi/govc/device/cdrom:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/device/floppy:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/device/scsi:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/device/serial:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/device/usb:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
