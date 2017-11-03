load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "environment.go",
        "experiment.go",
        "group.go",
        "mount.go",
        "resolv_conf.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgroup_util.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/common",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fs:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/group:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/mountinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "resolv_conf_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgroup_util_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/common",
    library = ":go_default_library",
    deps = ["//vendor/github.com/containernetworking/cni/pkg/types:go_default_library"],
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
        "//vendor/github.com/coreos/rkt/common/apps:all-srcs",
        "//vendor/github.com/coreos/rkt/common/cgroup:all-srcs",
        "//vendor/github.com/coreos/rkt/common/labelsort:all-srcs",
        "//vendor/github.com/coreos/rkt/common/networking:all-srcs",
        "//vendor/github.com/coreos/rkt/common/overlay:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
