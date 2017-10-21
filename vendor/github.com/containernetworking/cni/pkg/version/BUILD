load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "conf.go",
        "plugin.go",
        "reconcile.go",
        "version.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/version",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/020:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/current:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "conf_test.go",
        "plugin_test.go",
        "reconcile_test.go",
        "version_suite_test.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/version_test",
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
