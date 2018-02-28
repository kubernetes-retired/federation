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
        "//vendor/github.com/influxdata/influxdb/cmd/influx_tsm/b1:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_tsm/bz1:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_tsm/stats:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_tsm/tsdb:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
