package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "types.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/proxy",
    deps = ["//vendor/k8s.io/apimachinery/pkg/types:go_default_library"],
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
        "//vendor/k8s.io/kubernetes/pkg/proxy/config:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/healthcheck:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/iptables:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/ipvs:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/userspace:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/util:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/winkernel:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/proxy/winuserspace:all-srcs",
    ],
    tags = ["automanaged"],
)
