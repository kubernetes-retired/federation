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
        "//vendor/github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/awstesting/mock:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/awstesting/unit:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
