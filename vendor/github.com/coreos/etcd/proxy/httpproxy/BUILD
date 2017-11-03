load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "director.go",
        "doc.go",
        "metrics.go",
        "proxy.go",
        "reverse.go",
    ],
    importpath = "github.com/coreos/etcd/proxy/httpproxy",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/etcd/etcdserver/api/v2http/httptypes:go_default_library",
        "//vendor/github.com/coreos/etcd/pkg/httputil:go_default_library",
        "//vendor/github.com/coreos/pkg/capnslog:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/golang.org/x/net/http2:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "director_test.go",
        "proxy_test.go",
        "reverse_test.go",
    ],
    importpath = "github.com/coreos/etcd/proxy/httpproxy",
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
