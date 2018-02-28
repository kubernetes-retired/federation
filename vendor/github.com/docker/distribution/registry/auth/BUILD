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
        "//vendor/github.com/docker/distribution/registry/auth/htpasswd:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth/silly:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth/token:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
