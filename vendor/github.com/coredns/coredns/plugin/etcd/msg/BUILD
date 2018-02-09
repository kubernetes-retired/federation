load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "path.go",
        "service.go",
        "type.go",
    ],
    importpath = "github.com/coredns/coredns/plugin/etcd/msg",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coredns/coredns/plugin/pkg/dnsutil:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "path_test.go",
        "service_test.go",
        "type_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/coredns/coredns/plugin/etcd/msg",
    deps = ["//vendor/github.com/miekg/dns:go_default_library"],
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
