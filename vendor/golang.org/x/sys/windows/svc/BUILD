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
        "//vendor/golang.org/x/sys/windows/svc/debug:all-srcs",
        "//vendor/golang.org/x/sys/windows/svc/eventlog:all-srcs",
        "//vendor/golang.org/x/sys/windows/svc/example:all-srcs",
        "//vendor/golang.org/x/sys/windows/svc/mgr:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
