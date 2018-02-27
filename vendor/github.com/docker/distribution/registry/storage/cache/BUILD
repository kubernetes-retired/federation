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
        "//vendor/github.com/docker/distribution/registry/storage/cache/cachecheck:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/cache/memory:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage/cache/redis:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
