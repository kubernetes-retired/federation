load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["labels.go"],
    importpath = "github.com/docker/libnetwork/netlabel",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["labels_test.go"],
    importpath = "github.com/docker/libnetwork/netlabel",
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
