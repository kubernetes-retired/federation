load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["request.go"],
    importpath = "github.com/miekg/coredns/request",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/miekg/coredns/middleware/pkg/edns:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["request_test.go"],
    importpath = "github.com/miekg/coredns/request",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/miekg/coredns/middleware/test:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
