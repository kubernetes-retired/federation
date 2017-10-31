load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cname.go",
        "dedup.go",
        "doc.go",
        "reverse.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/pkg/dnsutil",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/miekg/dns:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cname_test.go",
        "reverse_test.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/pkg/dnsutil",
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
