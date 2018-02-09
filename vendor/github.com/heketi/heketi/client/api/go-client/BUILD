load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "backup.go",
        "client.go",
        "cluster.go",
        "device.go",
        "node.go",
        "topology.go",
        "volume.go",
    ],
    importpath = "github.com/heketi/heketi/client/api/go-client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["client_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/heketi/heketi/client/api/go-client",
    deps = [
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/heketi/heketi/apps/glusterfs:go_default_library",
        "//vendor/github.com/heketi/heketi/middleware:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
        "//vendor/github.com/urfave/negroni:go_default_library",
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
