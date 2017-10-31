load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "args.go",
        "types.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/types",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "args_test.go",
        "types_suite_test.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/types_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/ginkgo/extensions/table:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
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
        "//vendor/github.com/containernetworking/cni/pkg/types/020:all-srcs",
        "//vendor/github.com/containernetworking/cni/pkg/types/current:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
