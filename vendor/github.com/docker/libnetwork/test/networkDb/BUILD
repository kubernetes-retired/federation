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
        "//vendor/github.com/docker/libnetwork/test/networkDb/dbclient:all-srcs",
        "//vendor/github.com/docker/libnetwork/test/networkDb/dbserver:all-srcs",
        "//vendor/github.com/docker/libnetwork/test/networkDb/dummyclient:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
