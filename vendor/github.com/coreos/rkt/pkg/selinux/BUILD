load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "selinux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/pkg/selinux",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_xtest",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "selinux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/rkt/pkg/selinux_test",
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            ":go_default_library",
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
