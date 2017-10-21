package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = ["providers.go"],
    importpath = "k8s.io/kubernetes/pkg/cloudprovider/providers",
    deps = [
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/aws:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/azure:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/cloudstack:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/gce:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/openstack:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/ovirt:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/photon:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere:go_default_library",
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
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/aws:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/azure:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/cloudstack:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/gce:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/openstack:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/ovirt:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/photon:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/cloudprovider/providers/vsphere:all-srcs",
    ],
    tags = ["automanaged"],
)
