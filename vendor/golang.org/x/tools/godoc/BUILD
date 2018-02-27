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
        "//vendor/golang.org/x/tools/godoc/analysis:all-srcs",
        "//vendor/golang.org/x/tools/godoc/redirect:all-srcs",
        "//vendor/golang.org/x/tools/godoc/static:all-srcs",
        "//vendor/golang.org/x/tools/godoc/util:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
