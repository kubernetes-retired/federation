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
        "//vendor/github.com/docker/distribution/registry/api/errcode:all-srcs",
        "//vendor/github.com/docker/distribution/registry/api/v2:all-srcs",
        "//vendor/github.com/docker/distribution/registry/auth:all-srcs",
        "//vendor/github.com/docker/distribution/registry/client:all-srcs",
        "//vendor/github.com/docker/distribution/registry/handlers:all-srcs",
        "//vendor/github.com/docker/distribution/registry/listener:all-srcs",
        "//vendor/github.com/docker/distribution/registry/middleware/registry:all-srcs",
        "//vendor/github.com/docker/distribution/registry/middleware/repository:all-srcs",
        "//vendor/github.com/docker/distribution/registry/proxy:all-srcs",
        "//vendor/github.com/docker/distribution/registry/storage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
