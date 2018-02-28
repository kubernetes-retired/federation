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
        "//vendor/github.com/docker/docker/daemon/caps:all-srcs",
        "//vendor/github.com/docker/docker/daemon/cluster:all-srcs",
        "//vendor/github.com/docker/docker/daemon/config:all-srcs",
        "//vendor/github.com/docker/docker/daemon/discovery:all-srcs",
        "//vendor/github.com/docker/docker/daemon/events:all-srcs",
        "//vendor/github.com/docker/docker/daemon/exec:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver:all-srcs",
        "//vendor/github.com/docker/docker/daemon/initlayer:all-srcs",
        "//vendor/github.com/docker/docker/daemon/links:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger:all-srcs",
        "//vendor/github.com/docker/docker/daemon/network:all-srcs",
        "//vendor/github.com/docker/docker/daemon/stats:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
