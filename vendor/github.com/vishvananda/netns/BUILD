load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "netns.go",
        "netns_unspecified.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "netns_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/vishvananda/netns",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["netns_test.go"],
    importpath = "github.com/vishvananda/netns",
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
