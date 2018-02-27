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
        "//vendor/github.com/docker/docker/pkg/plugins/pluginrpc-gen:all-srcs",
        "//vendor/github.com/docker/docker/pkg/plugins/transport:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
