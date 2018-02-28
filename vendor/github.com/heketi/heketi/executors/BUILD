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
        "//vendor/github.com/heketi/heketi/executors/kubeexec:all-srcs",
        "//vendor/github.com/heketi/heketi/executors/mockexec:all-srcs",
        "//vendor/github.com/heketi/heketi/executors/sshexec:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
