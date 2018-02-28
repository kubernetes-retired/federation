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
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect/dumptsm:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect/export:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect/help:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect/report:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect/verify:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
