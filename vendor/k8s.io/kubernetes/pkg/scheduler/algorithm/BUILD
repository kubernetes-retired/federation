load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "scheduler_interface.go",
        "types.go",
        "well_known_labels.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/scheduler/algorithm",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/apps/v1beta1:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/api/extensions/v1beta1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/scheduler/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/scheduler/schedulercache:go_default_library",
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
        "//vendor/k8s.io/kubernetes/pkg/scheduler/algorithm/predicates:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/scheduler/algorithm/priorities/util:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
