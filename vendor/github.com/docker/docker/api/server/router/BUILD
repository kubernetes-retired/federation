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
        "//vendor/github.com/docker/docker/api/server/router/build:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/checkpoint:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/container:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/debug:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/distribution:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/image:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/network:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/plugin:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/session:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/swarm:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/system:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router/volume:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
