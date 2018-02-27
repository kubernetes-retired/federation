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
        "//vendor/github.com/gogo/protobuf/vanity/command:all-srcs",
        "//vendor/github.com/gogo/protobuf/vanity/test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
