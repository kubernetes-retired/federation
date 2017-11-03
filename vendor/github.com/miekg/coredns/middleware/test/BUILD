load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "file.go",
        "helpers.go",
        "responsewriter.go",
        "server.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/test",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/miekg/dns:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "file_test.go",
        "helpers_test.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/test",
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
