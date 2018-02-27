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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze/statistics:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze/summarize:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
