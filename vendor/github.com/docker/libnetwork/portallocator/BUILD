load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "portallocator.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "portallocator_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/portallocator",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["portallocator_test.go"],
    importpath = "github.com/docker/libnetwork/portallocator",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/libnetwork/testutils:go_default_library"],
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
