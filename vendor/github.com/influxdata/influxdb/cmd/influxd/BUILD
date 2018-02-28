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
        "//vendor/github.com/influxdata/influxdb/cmd/influxd/backup:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influxd/help:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influxd/restore:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influxd/run:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
