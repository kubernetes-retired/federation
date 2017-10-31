load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["flag.go"],
    importpath = "github.com/docker/libnetwork/client/mflag",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/homedir:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["flag_test.go"],
    importpath = "github.com/docker/libnetwork/client/mflag",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/libnetwork/client/mflag/example:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
