load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "gotype.go",
        "sizesFor18.go",
        "sizesFor19.go",
    ],
    importpath = "golang.org/x/tools/cmd/gotype",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "gotype",
    importpath = "golang.org/x/tools/cmd/gotype",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
