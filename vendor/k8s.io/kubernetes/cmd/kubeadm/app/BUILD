package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["kubeadm.go"],
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app",
    deps = [
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/install:go_default_library",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd:go_default_library",
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
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util:all-srcs",
    ],
    tags = ["automanaged"],
)
