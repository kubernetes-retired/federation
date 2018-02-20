load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["e2e.go"],
    importpath = "k8s.io/kubernetes/hack",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "hack",
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/hack",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["e2e_test.go"],
    data = glob(["testdata/**"]),
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/hack",
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
        "//vendor/k8s.io/kubernetes/hack/boilerplate:all-srcs",
        "//vendor/k8s.io/kubernetes/hack/cmd/teststale:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
