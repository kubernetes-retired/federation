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
        "//vendor/github.com/docker/docker/daemon/graphdriver/aufs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/btrfs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/devmapper:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/graphtest:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/lcow:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/overlay:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/overlay2:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/overlayutils:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/quota:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/register:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/vfs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/windows:all-srcs",
        "//vendor/github.com/docker/docker/daemon/graphdriver/zfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
