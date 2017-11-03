load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["connection.go"],
    importpath = "github.com/docker/spdystream/ws",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/gorilla/websocket:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["ws_test.go"],
    importpath = "github.com/docker/spdystream/ws",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/spdystream:go_default_library",
        "//vendor/github.com/gorilla/websocket:go_default_library",
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
