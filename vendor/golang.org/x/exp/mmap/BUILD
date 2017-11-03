load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "mmap_other.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "mmap_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "mmap_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/mmap",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["mmap_test.go"],
    importpath = "golang.org/x/exp/mmap",
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
