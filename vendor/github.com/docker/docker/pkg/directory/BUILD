load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "directory.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "directory_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "directory_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/pkg/directory",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["directory_test.go"],
    importpath = "github.com/docker/docker/pkg/directory",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
