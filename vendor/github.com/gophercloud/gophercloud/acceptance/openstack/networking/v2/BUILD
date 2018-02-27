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
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
