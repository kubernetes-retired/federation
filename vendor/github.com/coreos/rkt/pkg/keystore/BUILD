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
        "//vendor/github.com/coreos/rkt/pkg/keystore/keystoretest:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
