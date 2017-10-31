load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "gc.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/stage1/gc",
    visibility = ["//visibility:private"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/appc/spec/schema/types:go_default_library",
            "//vendor/github.com/coreos/rkt/common:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup:go_default_library",
            "//vendor/github.com/coreos/rkt/common/cgroup/v1:go_default_library",
            "//vendor/github.com/coreos/rkt/networking:go_default_library",
            "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
            "//vendor/github.com/hashicorp/errwrap:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_binary(
    name = "gc",
    importpath = "github.com/coreos/rkt/stage1/gc",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
