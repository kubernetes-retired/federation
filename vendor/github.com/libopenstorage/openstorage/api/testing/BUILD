load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["testing.go"],
    importpath = "github.com/libopenstorage/openstorage/api/testing",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["testing_test.go"],
    importpath = "github.com/libopenstorage/openstorage/api/testing",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/client/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/server:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/nfs:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/test:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
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
