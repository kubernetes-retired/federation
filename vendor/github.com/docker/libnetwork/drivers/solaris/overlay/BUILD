load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["overlay.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "encryption.go",
        "joinleave.go",
        "ov_endpoint.go",
        "ov_network.go",
        "ov_serf.go",
        "ov_utils.go",
        "overlay.go",
        "overlay.pb.go",
        "peerdb.go",
    ],
    importpath = "github.com/docker/libnetwork/drivers/solaris/overlay",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/libnetwork/datastore:go_default_library",
        "//vendor/github.com/docker/libnetwork/discoverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/driverapi:go_default_library",
        "//vendor/github.com/docker/libnetwork/idm:go_default_library",
        "//vendor/github.com/docker/libnetwork/netlabel:go_default_library",
        "//vendor/github.com/docker/libnetwork/netutils:go_default_library",
        "//vendor/github.com/docker/libnetwork/osl:go_default_library",
        "//vendor/github.com/docker/libnetwork/resolvconf:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/hashicorp/serf/serf:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
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
        "//vendor/github.com/docker/libnetwork/drivers/solaris/overlay/ovmanager:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
