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
        "//vendor/github.com/docker/docker/pkg/discovery/file:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/kv:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/memory:all-srcs",
        "//vendor/github.com/docker/docker/pkg/discovery/nodes:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
