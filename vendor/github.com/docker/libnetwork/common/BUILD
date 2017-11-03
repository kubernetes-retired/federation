load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "caller.go",
        "setmatrix.go",
    ],
    importpath = "github.com/docker/libnetwork/common",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/deckarep/golang-set:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "caller_test.go",
        "setmatrix_test.go",
    ],
    importpath = "github.com/docker/libnetwork/common",
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
