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
        "//vendor/golang.org/x/tools/godoc/vfs/gatefs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/httpfs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/mapfs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/zipfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
