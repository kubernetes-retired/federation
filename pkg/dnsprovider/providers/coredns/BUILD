package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "coredns.go",
        "interface.go",
        "rrchangeset.go",
        "rrset.go",
        "rrsets.go",
        "zone.go",
        "zones.go",
    ],
    importpath = "k8s.io/federation/pkg/dnsprovider/providers/coredns",
    deps = [
        "//pkg/dnsprovider:go_default_library",
        "//pkg/dnsprovider/providers/coredns/stubs:go_default_library",
        "//pkg/dnsprovider/rrstype:go_default_library",
        "//vendor/github.com/coreos/etcd/client:go_default_library",
        "//vendor/github.com/coreos/etcd/pkg/transport:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/etcd/msg:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/gopkg.in/gcfg.v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["coredns_test.go"],
    importpath = "k8s.io/federation/pkg/dnsprovider/providers/coredns",
    library = ":go_default_library",
    deps = [
        "//pkg/dnsprovider:go_default_library",
        "//pkg/dnsprovider/providers/coredns/stubs:go_default_library",
        "//pkg/dnsprovider/rrstype:go_default_library",
        "//pkg/dnsprovider/tests:go_default_library",
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
        "//pkg/dnsprovider/providers/coredns/stubs:all-srcs",
    ],
    tags = ["automanaged"],
)
