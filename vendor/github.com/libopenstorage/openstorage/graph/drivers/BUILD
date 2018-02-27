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
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/chainfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/layer0:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/graph/drivers/proxy:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
