load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "network.go",
        "service.go",
        "types.go",
    ],
    importpath = "github.com/docker/libnetwork/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/opts:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/docker/libnetwork/client/mflag:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_service_test.go",
        "client_test.go",
    ],
    importpath = "github.com/docker/libnetwork/client",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/libnetwork/testutils:go_default_library"],
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
        "//vendor/github.com/docker/libnetwork/client/mflag:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
