load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "example_runtime.go",
        "examples.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/version/legacy_examples",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/020:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version/testhelpers:go_default_library",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "legacy_examples_suite_test.go",
        "legacy_examples_test.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/version/legacy_examples_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
