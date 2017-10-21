load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cgroups_unsupported.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgroups.go",
            "stats.go",
            "utils.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/opencontainers/runc/libcontainer/cgroups",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/docker/go-units:go_default_library",
            "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "cgroups_test.go",
            "utils_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/opencontainers/runc/libcontainer/cgroups",
    library = ":go_default_library",
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
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups/fs:all-srcs",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups/rootless:all-srcs",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups/systemd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
