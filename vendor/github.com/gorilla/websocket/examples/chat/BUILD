load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "conn.go",
        "hub.go",
        "main.go",
    ],
    importpath = "github.com/gorilla/websocket/examples/chat",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/gorilla/websocket:go_default_library"],
)

go_binary(
    name = "chat",
    importpath = "github.com/gorilla/websocket/examples/chat",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
