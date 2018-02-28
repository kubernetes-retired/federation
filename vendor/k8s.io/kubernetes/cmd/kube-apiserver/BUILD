package(default_visibility = ["//visibility:public"])

load("//pkg/version:def.bzl", "version_x_defs")

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
        "//vendor/k8s.io/kubernetes/cmd/kube-apiserver/app:all-srcs",
    ],
    tags = ["automanaged"],
)
