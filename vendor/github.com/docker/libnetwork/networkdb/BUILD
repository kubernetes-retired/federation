load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = ["networkdb.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "broadcast.go",
        "cluster.go",
        "delegate.go",
        "event_delegate.go",
        "message.go",
        "networkdb.go",
        "networkdb.pb.go",
        "networkdbdiagnose.go",
        "watch.go",
    ],
    importpath = "github.com/docker/libnetwork/networkdb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/armon/go-radix:go_default_library",
        "//vendor/github.com/docker/go-events:go_default_library",
        "//vendor/github.com/docker/libnetwork/diagnose:go_default_library",
        "//vendor/github.com/docker/libnetwork/types:go_default_library",
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/hashicorp/memberlist:go_default_library",
        "//vendor/github.com/hashicorp/serf/serf:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["networkdb_test.go"],
    importpath = "github.com/docker/libnetwork/networkdb",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/go-events:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
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
