load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "kubeexec.go",
    ],
    importpath = "github.com/heketi/heketi/executors/kubeexec",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/heketi/heketi/executors/sshexec:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/kubernetes:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/lpabon/godbc:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset/typed/core/v1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/unversioned/remotecommand:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/server/remotecommand:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["kubeexec_test.go"],
    importpath = "github.com/heketi/heketi/executors/kubeexec",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/heketi/heketi/executors/sshexec:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
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
