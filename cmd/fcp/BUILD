load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "federation-apiserver.go",
        "federation-controller-manager.go",
        "hyperkube.go",
        "main.go",
        "server.go",
    ],
    importpath = "k8s.io/federation/cmd/fcp",
    visibility = ["//visibility:private"],
    deps = [
        "//cmd/federation-apiserver/app:go_default_library",
        "//cmd/federation-apiserver/app/options:go_default_library",
        "//cmd/federation-controller-manager/app:go_default_library",
        "//cmd/federation-controller-manager/app/options:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/metrics/prometheus:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/template:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version/prometheus:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version/verflag:go_default_library",
    ],
)

go_binary(
    name = "fcp",
    importpath = "k8s.io/federation/cmd/fcp",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
