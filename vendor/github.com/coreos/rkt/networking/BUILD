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
        "//vendor/github.com/coreos/rkt/networking/netinfo:all-srcs",
        "//vendor/github.com/coreos/rkt/networking/tuntap:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
