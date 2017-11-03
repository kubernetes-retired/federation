load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "capability.go",
        "sys.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "sys_linux.go",
            "sys_linux_amd64.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/pkg/sys",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/syndtr/gocapability/capability:go_default_library"],
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
