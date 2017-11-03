load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cli.go",
        "cluster.go",
        "drivers.go",
        "graph.go",
        "messages.go",
        "volumes.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/cli",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/codegangsta/cli:go_default_library",
        "//vendor/github.com/dustin/go-humanize:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/libopenstorage/gossip/types:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/client/cluster:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api/client/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/cluster:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/jsonpb:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["cli_test.go"],
    importpath = "github.com/libopenstorage/openstorage/cli",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
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
