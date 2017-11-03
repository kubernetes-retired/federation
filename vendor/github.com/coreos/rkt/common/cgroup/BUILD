load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgroup.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/common/cgroup",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/coreos/rkt/common/cgroup/v1:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup/v2:go_default_library",
            "//vendor/github.com/hashicorp/errwrap:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

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
        "//vendor/github.com/coreos/rkt/common/cgroup/v1:all-srcs",
        "//vendor/github.com/coreos/rkt/common/cgroup/v2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
