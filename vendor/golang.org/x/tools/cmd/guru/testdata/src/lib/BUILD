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
        "//vendor/golang.org/x/tools/cmd/guru/testdata/src/lib/sublib:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
