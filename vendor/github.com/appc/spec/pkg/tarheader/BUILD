load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "tarheader.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "pop_darwin.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "pop_linux.go",
            "pop_posix.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/appc/spec/pkg/tarheader",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/appc/spec/pkg/device:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "pop_posix_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/appc/spec/pkg/tarheader",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
