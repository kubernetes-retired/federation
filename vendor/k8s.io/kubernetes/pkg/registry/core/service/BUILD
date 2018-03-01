load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "proxy.go",
        "registry.go",
        "rest.go",
        "strategy.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/registry/core/service",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/net:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/proxy:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/validation/field:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/endpoints/request:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/registry/rest:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/storage/names:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api/legacyscheme:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api/service:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core/helper:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core/validation:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/capabilities:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/features:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/endpoint:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/pod/storage:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/ipallocator:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/portallocator:go_default_library",
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
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/allocator:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/ipallocator:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/portallocator:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/registry/core/service/storage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
