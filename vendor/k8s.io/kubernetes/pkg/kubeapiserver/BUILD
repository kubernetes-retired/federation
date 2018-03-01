load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "default_storage_factory_builder.go",
        "doc.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/kubeapiserver",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server/resourceconfig:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server/storage:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/storage/storagebackend:go_default_library",
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
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/admission:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/authenticator:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/authorizer:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/options:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/server:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
