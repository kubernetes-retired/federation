load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["mapfs.go"],
    importpath = "golang.org/x/tools/godoc/vfs/mapfs",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/godoc/vfs:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["mapfs_test.go"],
    importpath = "golang.org/x/tools/godoc/vfs/mapfs",
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
