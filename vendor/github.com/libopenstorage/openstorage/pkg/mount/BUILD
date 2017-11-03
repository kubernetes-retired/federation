load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "custom_mount.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "device.go",
            "mount.go",
            "nfs.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/libopenstorage/openstorage/pkg/mount",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/docker/pkg/mount:go_default_library",
            "//vendor/go.pedge.io/dlog:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = ["mount_test.go"],
    importpath = "github.com/libopenstorage/openstorage/pkg/mount",
    library = ":go_default_library",
    deps = ["//vendor/github.com/stretchr/testify/require:go_default_library"],
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
