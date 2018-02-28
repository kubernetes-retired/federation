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
        "//vendor/github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3/s3manager:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
