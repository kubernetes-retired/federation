load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "ripemd160.go",
        "ripemd160block.go",
    ],
    importpath = "golang.org/x/crypto/ripemd160",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["ripemd160_test.go"],
    importpath = "golang.org/x/crypto/ripemd160",
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
