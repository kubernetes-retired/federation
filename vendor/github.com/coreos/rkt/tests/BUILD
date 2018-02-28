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
        "//vendor/github.com/coreos/rkt/tests/cniproxy:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/echo-socket-activated:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/inspect:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/enter:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/gc:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/run:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/testutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
