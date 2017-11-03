load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["memcache.go"],
    importpath = "github.com/gregjones/httpcache/memcache",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/bradfitz/gomemcache/memcache:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["memcache_test.go"],
    importpath = "github.com/gregjones/httpcache/memcache",
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
