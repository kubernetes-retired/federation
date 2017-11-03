load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "conn.go",
        "doc.go",
        "json.go",
        "server.go",
        "util.go",
    ],
    importpath = "github.com/gorilla/websocket",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bench_test.go",
        "client_server_test.go",
        "client_test.go",
        "conn_test.go",
        "json_test.go",
        "server_test.go",
        "util_test.go",
    ],
    importpath = "github.com/gorilla/websocket",
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
        "//vendor/github.com/gorilla/websocket/examples/autobahn:all-srcs",
        "//vendor/github.com/gorilla/websocket/examples/chat:all-srcs",
        "//vendor/github.com/gorilla/websocket/examples/filewatch:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
