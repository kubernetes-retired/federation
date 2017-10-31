load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_test(
    name = "go_default_test",
    srcs = ["bookstore_test.go"],
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore",
    deps = ["//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore/bookstore:go_default_library"],
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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore/service:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
