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
        "//vendor/github.com/miekg/coredns/middleware/metrics/test:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/metrics/vars:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
