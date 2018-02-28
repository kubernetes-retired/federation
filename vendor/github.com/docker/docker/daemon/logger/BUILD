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
        "//vendor/github.com/docker/docker/daemon/logger/awslogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/etwlogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/fluentd:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/gcplogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/gelf:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/journald:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/jsonfilelog:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/logentries:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/loggerutils:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/splunk:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/syslog:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
