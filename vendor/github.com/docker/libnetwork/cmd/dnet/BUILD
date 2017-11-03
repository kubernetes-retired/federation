load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cmd.go",
        "dnet.go",
        "flags.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "dnet_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "dnet_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/cmd/dnet",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/codegangsta/cli:go_default_library",
        "//vendor/github.com/docker/docker/api/types/network:go_default_library",
        "//vendor/github.com/docker/docker/opts:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/docker/libnetwork:go_default_library",
        "//vendor/github.com/docker/libnetwork/api:go_default_library",
        "//vendor/github.com/docker/libnetwork/client:go_default_library",
        "//vendor/github.com/docker/libnetwork/cluster:go_default_library",
        "//vendor/github.com/docker/libnetwork/config:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/options:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/docker/pkg/signal:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/docker/docker/pkg/signal:go_default_library",
            "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_binary(
    name = "dnet",
    importpath = "github.com/docker/libnetwork/cmd/dnet",
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
