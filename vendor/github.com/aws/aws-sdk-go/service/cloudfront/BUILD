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
        "//vendor/github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/service/cloudfront/sign:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
