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
        "//vendor/github.com/docker/distribution/registry/storage/driver/middleware/cloudfront:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/driver/middleware/redirect:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
