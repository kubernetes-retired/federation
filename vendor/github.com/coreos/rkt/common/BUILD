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
        "//vendor/github.com/coreos/rkt/common/apps:all-srcs",
        "//vendor/github.com/coreos/rkt/common/cgroup:all-srcs",
        "//vendor/github.com/coreos/rkt/common/labelsort:all-srcs",
        "//vendor/github.com/coreos/rkt/common/networking:all-srcs",
        "//vendor/github.com/coreos/rkt/common/overlay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
