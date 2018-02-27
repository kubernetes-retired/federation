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
        "//vendor/go4.org/syncutil/singleflight:all-srcs",
        "//vendor/go4.org/syncutil/syncdebug:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
