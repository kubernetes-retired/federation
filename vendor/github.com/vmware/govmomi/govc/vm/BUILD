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
        "//vendor/github.com/vmware/govmomi/govc/vm/disk:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/vm/guest:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/vm/network:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/vm/rdm:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/vm/snapshot:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
