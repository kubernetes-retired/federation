load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "common_unix.go",
        "names.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "common_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        "//vendor/github.com/docker/libtrust:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["common_test.go"],
    importpath = "github.com/docker/docker/api",
    library = ":go_default_library",
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
        "//vendor/github.com/docker/docker/api/errors:all-srcs",
        "//vendor/github.com/docker/docker/api/server:all-srcs",
        "//vendor/github.com/docker/docker/api/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
