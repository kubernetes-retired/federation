load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "merge.go",
        "merge_visitor.go",
        "replace_visitor.go",
        "retain_keys_visitor.go",
        "strategic_visitor.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubectl/apply/strategy",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/kubernetes/pkg/kubectl/apply:go_default_library"],
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
