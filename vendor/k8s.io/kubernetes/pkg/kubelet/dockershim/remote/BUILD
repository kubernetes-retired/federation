load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["docker_server.go"],
    importpath = "k8s.io/kubernetes/pkg/kubelet/dockershim/remote",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/dockershim:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/util:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/interrupt:go_default_library",
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
