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
        "//vendor/github.com/docker/docker/api/server/backend/build:all-srcs",
        "//vendor/github.com/docker/docker/api/server/httputils:all-srcs",
        "//vendor/github.com/docker/docker/api/server/middleware:all-srcs",
        "//vendor/github.com/docker/docker/api/server/router:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
