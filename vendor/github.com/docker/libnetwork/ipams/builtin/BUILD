load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "builtin_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "builtin_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "builtin_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/ipams/builtin",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipam:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipam:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipam:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipams/windowsipam:go_default_library",
            "//vendor/github.com/docker/libnetwork/ipamutils:go_default_library",
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
