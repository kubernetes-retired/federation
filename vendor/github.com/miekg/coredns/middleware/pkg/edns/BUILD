load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["edns.go"],
    importpath = "github.com/miekg/coredns/middleware/pkg/edns",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/miekg/dns:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["edns_test.go"],
    importpath = "github.com/miekg/coredns/middleware/pkg/edns",
    library = ":go_default_library",
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
