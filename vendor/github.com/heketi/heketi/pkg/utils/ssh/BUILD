load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["ssh.go"],
    importpath = "github.com/heketi/heketi/pkg/utils/ssh",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/agent:go_default_library",
    ],
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
