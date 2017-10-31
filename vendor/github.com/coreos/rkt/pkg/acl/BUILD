load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "acl.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "-ldl",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/pkg/acl",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
