load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["admission.go"],
    importpath = "k8s.io/kubernetes/plugin/pkg/admission/persistentvolume/resize",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apiserver/pkg/admission:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core/helper:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/listers/core/internalversion:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/listers/storage/internalversion:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubeapiserver/admission:go_default_library",
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
