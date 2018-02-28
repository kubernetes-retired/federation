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
        "//vendor/github.com/vmware/govmomi/govc/datastore/disk:all-srcs",
        "//vendor/github.com/vmware/govmomi/govc/datastore/vsan:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
