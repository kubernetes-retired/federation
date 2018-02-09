load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "dial.go",
        "hybi.go",
        "server.go",
        "websocket.go",
    ],
    importpath = "golang.org/x/net/websocket",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "dial_test.go",
        "hybi_test.go",
        "websocket_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "golang.org/x/net/websocket",
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "exampledial_test.go",
        "examplehandler_test.go",
    ],
    importpath = "golang.org/x/net/websocket_test",
    deps = [":go_default_library"],
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
