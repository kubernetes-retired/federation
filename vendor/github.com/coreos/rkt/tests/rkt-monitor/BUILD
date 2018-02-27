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
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/cpu-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/log-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/mem-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/sleeper:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
