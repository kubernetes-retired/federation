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
        "//vendor/github.com/coreos/etcd/tools/functional-tester/etcd-runner/command:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
