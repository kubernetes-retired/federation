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
        "//vendor/github.com/docker/docker/cli/config:all-srcs",
        "//vendor/github.com/docker/docker/cli/debug:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
