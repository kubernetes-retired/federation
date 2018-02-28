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
        "//vendor/github.com/aws/aws-sdk-go/service/elbv2/elbv2iface:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
