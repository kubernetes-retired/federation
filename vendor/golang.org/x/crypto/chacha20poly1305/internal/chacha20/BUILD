load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["chacha_generic.go"],
    importpath = "golang.org/x/crypto/chacha20poly1305/internal/chacha20",
    visibility = ["//vendor/golang.org/x/crypto/chacha20poly1305:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["chacha_test.go"],
    importpath = "golang.org/x/crypto/chacha20poly1305/internal/chacha20",
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
