load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "intelrdt.go",
            "stats.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/opencontainers/runc/libcontainer/intelrdt",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "intelrdt_test.go",
            "util_test.go",
        ],
        "//conditions:default": [],
    }),
    embed = [":go_default_library"],
    importpath = "github.com/opencontainers/runc/libcontainer/intelrdt",
    deps = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
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
