load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "brick.go",
        "config.go",
        "device.go",
        "peer.go",
        "sshexec.go",
        "volume.go",
    ],
    importpath = "github.com/heketi/heketi/executors/sshexec",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/heketi/heketi/executors:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils/ssh:go_default_library",
        "//vendor/github.com/lpabon/godbc:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "brick_test.go",
        "peer_test.go",
        "sshexec_test.go",
    ],
    importpath = "github.com/heketi/heketi/executors/sshexec",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/heketi/heketi/executors:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
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
