load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "netns.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "netns_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:nacl": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "netns_unspecified.go",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "netns_unspecified.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/vishvananda/netns",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["netns_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/vishvananda/netns",
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
