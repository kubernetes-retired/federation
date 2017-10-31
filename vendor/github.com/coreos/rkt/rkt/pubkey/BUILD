load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["pubkey.go"],
    importpath = "github.com/coreos/rkt/rkt/pubkey",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/discovery:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/keystore:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/rkt/config:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp:go_default_library",
        "//vendor/golang.org/x/crypto/ssh/terminal:go_default_library",
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
