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
        "//vendor/github.com/miekg/coredns/middleware/file/tree:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
