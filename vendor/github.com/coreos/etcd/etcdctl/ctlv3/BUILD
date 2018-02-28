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
        "//vendor/github.com/coreos/etcd/etcdctl/ctlv3/command:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
