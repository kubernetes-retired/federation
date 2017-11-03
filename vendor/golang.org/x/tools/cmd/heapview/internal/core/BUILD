load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "mmapfile_other.go",
        "raw.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "mmapfile.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "mmapfile.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/tools/cmd/heapview/internal/core",
    visibility = ["//vendor/golang.org/x/tools/cmd/heapview:__subpackages__"],
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
