load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "types.go",
    ],
    importpath = "github.com/docker/libnetwork/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "api_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "api_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/api",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/libnetwork/drivers/bridge:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
