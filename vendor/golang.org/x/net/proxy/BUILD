load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "direct.go",
        "per_host.go",
        "proxy.go",
        "socks5.go",
    ],
    importpath = "golang.org/x/net/proxy",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "per_host_test.go",
        "proxy_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "golang.org/x/net/proxy",
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
