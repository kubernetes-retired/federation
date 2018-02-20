load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "helpers.go",
        "util.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/apis/storage/util",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/storage:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/features:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["util_test.go"],
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/pkg/apis/storage/util",
    deps = [
        "//vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/storage:go_default_library",
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
