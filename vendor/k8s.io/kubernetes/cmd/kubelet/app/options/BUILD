load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "container_runtime.go",
        "globalflags.go",
        "options.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "globalflags_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:nacl": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "globalflags_other.go",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "globalflags_other.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "k8s.io/kubernetes/cmd/kubelet/app/options",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/feature:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/componentconfig:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/apis/core:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/azure:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/gcp:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/features:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/scheme:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/validation:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/config:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/types:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/taints:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version/verflag:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//vendor/github.com/google/cadvisor/container/common:go_default_library",
            "//vendor/github.com/google/cadvisor/container/containerd:go_default_library",
            "//vendor/github.com/google/cadvisor/container/docker:go_default_library",
            "//vendor/github.com/google/cadvisor/container/raw:go_default_library",
            "//vendor/github.com/google/cadvisor/machine:go_default_library",
            "//vendor/github.com/google/cadvisor/manager:go_default_library",
            "//vendor/github.com/google/cadvisor/storage:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = ["options_test.go"],
    embed = [":go_default_library"],
    importpath = "k8s.io/kubernetes/cmd/kubelet/app/options",
    deps = [
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/diff:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
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
