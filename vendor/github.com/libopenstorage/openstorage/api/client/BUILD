load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "request.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/api/client",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["request_test.go"],
    importpath = "github.com/libopenstorage/openstorage/api/client",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/libopenstorage/openstorage/api/client/cluster:all-srcs",
        "//vendor/github.com/libopenstorage/openstorage/api/client/volume:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
