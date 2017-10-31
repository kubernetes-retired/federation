load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "backend.go",
        "backend_lookup.go",
        "middleware.go",
        "normalize.go",
    ],
    importpath = "github.com/miekg/coredns/middleware",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/miekg/coredns/middleware/etcd/msg:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/pkg/dnsutil:go_default_library",
        "//vendor/github.com/miekg/coredns/request:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "middleware_test.go",
        "normalize_test.go",
    ],
    importpath = "github.com/miekg/coredns/middleware",
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
        "//vendor/github.com/miekg/coredns/middleware/auto:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/bind:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/cache:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/chaos:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/dnssec:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/errors:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/etcd:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/file:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/health:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/kubernetes:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/loadbalance:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/log:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/metrics:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/dnsrecorder:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/dnsutil:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/edns:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/rcode:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/replacer:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/response:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/singleflight:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/storage:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pkg/strings:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/pprof:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/proxy:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/rewrite:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/root:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/secondary:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/test:all-srcs",
        "//vendor/github.com/miekg/coredns/middleware/whoami:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
