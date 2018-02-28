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
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/aws:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/btrfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/buse:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/common:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/coprhd:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/fuse:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/nfs:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/pwx:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/test:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/vfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
