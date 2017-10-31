package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
)
load("//vendor/k8s.io/client-go/pkg/version:def.bzl", "version_x_defs")

go_binary(
    name = "kube-aggregator",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    importpath = "k8s.io/kube-aggregator",
    library = ":go_default_library",
    x_defs = version_x_defs(),
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "k8s.io/kube-aggregator",
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/install:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/validation:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/client/listers/apiregistration/v1beta1:go_default_library",
        "//vendor/k8s.io/kube-aggregator/pkg/cmd/server:go_default_library",
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
        "//vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/apiserver:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/informers/externalversions:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/informers/internalversion:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/listers/apiregistration/internalversion:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/client/listers/apiregistration/v1beta1:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/cmd/server:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/controllers:all-srcs",
        "//vendor/k8s.io/kube-aggregator/pkg/registry/apiservice:all-srcs",
    ],
    tags = ["automanaged"],
)
