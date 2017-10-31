load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["nametemplate.go"],
    importpath = "github.com/miekg/coredns/middleware/kubernetes/nametemplate",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/miekg/coredns/middleware/pkg/strings:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["nametemplate_test.go"],
    importpath = "github.com/miekg/coredns/middleware/kubernetes/nametemplate",
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
