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
        "//vendor/github.com/influxdata/influxdb/stress/v2/stressql/statement:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
