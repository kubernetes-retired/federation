load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["prototime.go"],
    importpath = "github.com/libopenstorage/openstorage/pkg/proto/time",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/ptypes/duration:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp:go_default_library",
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
