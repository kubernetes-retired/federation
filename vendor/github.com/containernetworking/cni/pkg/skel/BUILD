load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["skel.go"],
    importpath = "github.com/containernetworking/cni/pkg/skel",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "skel_suite_test.go",
        "skel_test.go",
    ],
    importpath = "github.com/containernetworking/cni/pkg/skel",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
