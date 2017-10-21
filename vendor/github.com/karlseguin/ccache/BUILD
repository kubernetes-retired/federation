load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bucket.go",
        "cache.go",
        "configuration.go",
        "item.go",
        "layeredbucket.go",
        "layeredcache.go",
        "secondarycache.go",
    ],
    importpath = "github.com/karlseguin/ccache",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bucket_test.go",
        "cache_test.go",
        "configuration_test.go",
        "item_test.go",
        "layeredcache_test.go",
        "secondarycache_test.go",
    ],
    importpath = "github.com/karlseguin/ccache",
    library = ":go_default_library",
    deps = ["//vendor/github.com/karlseguin/expect:go_default_library"],
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
