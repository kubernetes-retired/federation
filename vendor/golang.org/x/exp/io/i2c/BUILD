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
        "//vendor/golang.org/x/exp/io/i2c/driver:all-srcs",
        "//vendor/golang.org/x/exp/io/i2c/example/displayip:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
