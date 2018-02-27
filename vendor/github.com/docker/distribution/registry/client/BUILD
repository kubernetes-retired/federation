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
        "//vendor/github.com/docker/distribution/registry/client/auth:all-srcs",
        "//vendor/github.com/docker/distribution/registry/client/transport:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
