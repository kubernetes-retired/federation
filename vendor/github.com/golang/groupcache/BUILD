load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "byteview.go",
        "groupcache.go",
        "http.go",
        "peers.go",
        "sinks.go",
    ],
    importpath = "github.com/golang/groupcache",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/groupcache/consistenthash:go_default_library",
        "//vendor/github.com/golang/groupcache/groupcachepb:go_default_library",
        "//vendor/github.com/golang/groupcache/lru:go_default_library",
        "//vendor/github.com/golang/groupcache/singleflight:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "byteview_test.go",
        "groupcache_test.go",
        "http_test.go",
    ],
    importpath = "github.com/golang/groupcache",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/golang/groupcache/groupcachepb:go_default_library",
        "//vendor/github.com/golang/groupcache/testpb:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
    ],
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
        "//vendor/github.com/golang/groupcache/consistenthash:all-srcs",
        "//vendor/github.com/golang/groupcache/groupcachepb:all-srcs",
        "//vendor/github.com/golang/groupcache/lru:all-srcs",
        "//vendor/github.com/golang/groupcache/singleflight:all-srcs",
        "//vendor/github.com/golang/groupcache/testpb:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
