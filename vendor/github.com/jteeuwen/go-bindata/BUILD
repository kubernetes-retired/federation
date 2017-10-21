load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "asset.go",
        "bytewriter.go",
        "config.go",
        "convert.go",
        "debug.go",
        "doc.go",
        "release.go",
        "restore.go",
        "stringwriter.go",
        "toc.go",
    ],
    importpath = "github.com/jteeuwen/go-bindata",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["convert_test.go"],
    importpath = "github.com/jteeuwen/go-bindata",
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
        "//vendor/github.com/jteeuwen/go-bindata/go-bindata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
