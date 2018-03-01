load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["storage.go"],
    importpath = "k8s.io/kubernetes/pkg/registry/apps/statefulset/storage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/endpoints/request:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/registry/generic:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/registry/generic/registry:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/registry/rest:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/apps:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/apps/v1beta1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/apps/v1beta2:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/autoscaling:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/autoscaling/v1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/autoscaling/validation:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/extensions:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/printers:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/printers/internalversion:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/printers/storage:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/registry/apps/statefulset:go_default_library",
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
