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
