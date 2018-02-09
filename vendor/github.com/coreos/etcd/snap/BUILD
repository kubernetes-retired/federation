load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "db.go",
        "message.go",
        "metrics.go",
        "snapshotter.go",
    ],
    importpath = "github.com/coreos/etcd/snap",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/etcd/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/etcd/pkg/ioutil:go_default_library",
        "//vendor/github.com/coreos/etcd/pkg/pbutil:go_default_library",
        "//vendor/github.com/coreos/etcd/raft:go_default_library",
        "//vendor/github.com/coreos/etcd/raft/raftpb:go_default_library",
        "//vendor/github.com/coreos/etcd/snap/snappb:go_default_library",
        "//vendor/github.com/coreos/pkg/capnslog:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["snapshotter_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/coreos/etcd/snap",
    deps = ["//vendor/github.com/coreos/etcd/raft/raftpb:go_default_library"],
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
        "//vendor/github.com/coreos/etcd/snap/snappb:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
