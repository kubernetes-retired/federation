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
        "//vendor/golang.org/x/text/collate/build:all-srcs",
        "//vendor/golang.org/x/text/collate/tools/colcmp:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
