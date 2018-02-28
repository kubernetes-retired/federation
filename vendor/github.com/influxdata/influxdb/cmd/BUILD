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
        "//vendor/github.com/influxdata/influxdb/cmd/influx:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_inspect:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_stress:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influx_tsm:all-srcs",
        "//vendor/github.com/influxdata/influxdb/cmd/influxd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
