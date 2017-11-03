load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/google/cadvisor/storage/bigquery/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/golang.org/x/oauth2/jwt:go_default_library",
        "//vendor/google.golang.org/api/bigquery/v2:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/google/cadvisor/storage/bigquery/client/example:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
