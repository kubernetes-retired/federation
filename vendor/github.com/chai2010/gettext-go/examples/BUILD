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
        "//vendor/github.com/chai2010/gettext-go/examples/hi:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
