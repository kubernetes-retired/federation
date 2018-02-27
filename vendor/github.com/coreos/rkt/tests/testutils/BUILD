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
        "//vendor/github.com/coreos/rkt/tests/testutils/aci-server:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/testutils/logger:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
