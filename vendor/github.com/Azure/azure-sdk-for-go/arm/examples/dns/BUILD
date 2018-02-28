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
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/examples/dns/paging:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
