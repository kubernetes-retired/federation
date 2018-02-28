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
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/fwaas:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/layer3:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/lbaas:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/lbaas_v2:all-srcs",
        "//vendor/github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2/extensions/portsbinding:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
