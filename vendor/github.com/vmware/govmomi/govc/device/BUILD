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
