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
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/constants:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/discovery:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/features:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/images:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/dns:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/addons/proxy:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/clusterinfo:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/certs:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/controlplane:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/markmaster:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/selfhosting:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/token:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/upgrade:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/phases/uploadconfig:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/preflight:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util:all-srcs",
    ],
    tags = ["automanaged"],
)
