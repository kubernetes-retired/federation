load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "net.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "context_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "context_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/testutils",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/libnetwork/ns:go_default_library",
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
