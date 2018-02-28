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
        "//vendor/github.com/docker/docker/volume/drivers:all-srcs",
        "//vendor/github.com/docker/docker/volume/local:all-srcs",
        "//vendor/github.com/docker/docker/volume/store:all-srcs",
        "//vendor/github.com/docker/docker/volume/testutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
