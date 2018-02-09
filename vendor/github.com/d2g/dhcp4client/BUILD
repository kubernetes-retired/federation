load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "generatexid.go",
        "inetsock.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux": [
            "pktsock_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/d2g/dhcp4client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/d2g/dhcp4:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/golang.org/x/sys/unix:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "client_test.go",
        "generatexid_test.go",
    ],
    importpath = "github.com/d2g/dhcp4client_test",
    deps = [":go_default_library"],
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
