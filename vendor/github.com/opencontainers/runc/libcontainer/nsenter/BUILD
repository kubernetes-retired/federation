load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "namespace.h",
        "nsenter_unsupported.go",
        "nsexec.c",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "nsenter.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    copts = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "-Wall",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/opencontainers/runc/libcontainer/nsenter",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["nsenter_test.go"],
    importpath = "github.com/opencontainers/runc/libcontainer/nsenter",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/opencontainers/runc/libcontainer:go_default_library",
        "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
        "//vendor/golang.org/x/sys/unix:go_default_library",
    ],
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
