load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "conf.go",
    ],
    importpath = "github.com/containernetworking/cni/libcni",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/invoke:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "api_test.go",
        "backwards_compatibility_test.go",
        "conf_test.go",
        "libcni_suite_test.go",
    ],
    importpath = "github.com/containernetworking/cni/libcni_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/skel:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/current:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version/legacy_examples:go_default_library",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/onsi/gomega/gexec:go_default_library",
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
