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
        "//vendor/github.com/docker/docker/builder/dockerfile:all-srcs",
        "//vendor/github.com/docker/docker/builder/dockerignore:all-srcs",
        "//vendor/github.com/docker/docker/builder/fscache:all-srcs",
        "//vendor/github.com/docker/docker/builder/remotecontext:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
