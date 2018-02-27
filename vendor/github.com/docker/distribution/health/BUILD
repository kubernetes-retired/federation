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
        "//vendor/github.com/docker/distribution/health/api:all-srcs",
        "//vendor/github.com/docker/distribution/health/checks:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
