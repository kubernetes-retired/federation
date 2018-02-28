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
        "//vendor/github.com/docker/distribution/manifest/manifestlist:all-srcs",
        "//vendor/github.com/docker/distribution/manifest/schema1:all-srcs",
        "//vendor/github.com/docker/distribution/manifest/schema2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
