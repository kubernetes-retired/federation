load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["labelsort.go"],
    importpath = "github.com/coreos/rkt/common/labelsort",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/appc/spec/schema/types:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["labelsort_test.go"],
    importpath = "github.com/coreos/rkt/common/labelsort",
    library = ":go_default_library",
    deps = ["//vendor/github.com/appc/spec/schema/types:go_default_library"],
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
