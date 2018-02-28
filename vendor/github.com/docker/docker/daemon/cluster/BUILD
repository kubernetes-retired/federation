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
        "//vendor/github.com/docker/docker/daemon/cluster/controllers/plugin:all-srcs",
        "//vendor/github.com/docker/docker/daemon/cluster/convert:all-srcs",
        "//vendor/github.com/docker/docker/daemon/cluster/executor:all-srcs",
        "//vendor/github.com/docker/docker/daemon/cluster/provider:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
