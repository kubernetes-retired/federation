load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = ["overlay.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "encryption.go",
        "filter.go",
        "joinleave.go",
        "ostweaks_unsupported.go",
        "ov_endpoint.go",
        "ov_network.go",
        "ov_serf.go",
        "ov_utils.go",
        "overlay.go",
        "overlay.pb.go",
        "peerdb.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "ostweaks_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/libnetwork/drivers/overlay",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/libnetwork/common:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/idm:go_default_library",
        "//vendor/github.com/docker/libnetwork/iptables:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/ns:go_default_library",
        "//vendor/github.com/docker/libnetwork/osl:go_default_library",
        "//vendor/github.com/docker/libnetwork/resolvconf:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/hashicorp/serf/serf:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/vishvananda/netlink:go_default_library",
        "//vendor/github.com/vishvananda/netlink/nl:go_default_library",
        "//vendor/github.com/vishvananda/netns:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["overlay_test.go"],
    importpath = "github.com/docker/libnetwork/drivers/overlay",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/libkv/store/consul:go_default_library",
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/testutils:go_default_library",
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
        "//vendor/github.com/docker/libnetwork/drivers/overlay/ovmanager:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
