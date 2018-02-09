load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "grpclog.go",
        "logger.go",
        "loggerv2.go",
    ],
    importpath = "google.golang.org/grpc/grpclog",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["loggerv2_test.go"],
    embed = [":go_default_library"],
    importpath = "google.golang.org/grpc/grpclog",
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
