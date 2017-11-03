load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["overlay.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "overlay.pb.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "joinleave_windows.go",
            "ov_endpoint_windows.go",
            "ov_network_windows.go",
            "overlay_windows.go",
            "peerdb_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/drivers/windows/overlay",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Microsoft/hcsshim:go_default_library",
            "//vendor/github.com/docker/docker/pkg/system:go_default_library",
            "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
            "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
            "//vendor/github.com/docker/libnetwork/drivers/windows:go_default_library",
            "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
            "//vendor/github.com/docker/libnetwork/types:go_default_library",
            "//vendor/github.com/sirupsen/logrus:go_default_library",
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
