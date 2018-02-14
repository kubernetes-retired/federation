package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "dns.go",
        "doc.go",
        "plugins.go",
    ],
    importpath = "k8s.io/federation/pkg/dnsprovider",
    deps = [
        "//pkg/dnsprovider/rrstype:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["dns_test.go"],
    embed = [":go_default_library"],
    deps = ["//pkg/dnsprovider/rrstype:go_default_library"],
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
        "//pkg/dnsprovider/providers/aws/route53:all-srcs",
        "//pkg/dnsprovider/providers/coredns:all-srcs",
        "//pkg/dnsprovider/providers/google/clouddns:all-srcs",
        "//pkg/dnsprovider/rrstype:all-srcs",
        "//pkg/dnsprovider/tests:all-srcs",
    ],
    tags = ["automanaged"],
)
