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
        "//vendor/github.com/coreos/rkt/rkt/config:all-srcs",
        "//vendor/github.com/coreos/rkt/rkt/flag:all-srcs",
        "//vendor/github.com/coreos/rkt/rkt/image:all-srcs",
        "//vendor/github.com/coreos/rkt/rkt/pubkey:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
