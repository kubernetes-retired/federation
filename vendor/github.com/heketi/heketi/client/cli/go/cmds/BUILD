load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cluster.go",
        "device.go",
        "heketi_storage.go",
        "node.go",
        "root.go",
        "topology.go",
        "volume.go",
    ],
    importpath = "github.com/heketi/heketi/client/cli/go/cmds",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/heketi/heketi/client/api/go-client:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/db:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/kubernetes:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/batch/v1:go_default_library",
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
