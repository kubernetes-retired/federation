load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "stop_kvm.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/stage1/stop_kvm",
    visibility = ["//visibility:private"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
            "//vendor/github.com/coreos/rkt/stage1/common/ssh:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_binary(
    name = "stop_kvm",
    importpath = "github.com/coreos/rkt/stage1/stop_kvm",
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
