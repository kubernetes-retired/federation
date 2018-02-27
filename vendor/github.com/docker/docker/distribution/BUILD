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
        "//vendor/github.com/docker/docker/distribution/metadata:all-srcs",
        "//vendor/github.com/docker/docker/distribution/utils:all-srcs",
        "//vendor/github.com/docker/docker/distribution/xfer:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
