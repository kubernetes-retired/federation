load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["nnp.go"],
    importpath = "k8s.io/kubernetes/test/images/nonewprivs",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "nonewprivs",
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/test/images/nonewprivs",
    visibility = ["//visibility:public"],
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
