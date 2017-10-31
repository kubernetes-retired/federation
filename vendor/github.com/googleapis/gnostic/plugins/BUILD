load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["plugin.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["plugin.pb.go"],
    importpath = "github.com/googleapis/gnostic/plugins",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-sample:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
