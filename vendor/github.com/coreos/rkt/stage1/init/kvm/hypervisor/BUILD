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
        "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor/hvlkvm:all-srcs",
        "//vendor/github.com/coreos/rkt/stage1/init/kvm/hypervisor/hvqemu:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
