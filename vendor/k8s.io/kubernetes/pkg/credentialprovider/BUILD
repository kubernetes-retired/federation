package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "doc.go",
        "keyring.go",
        "plugins.go",
        "provider.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/credentialprovider",
    deps = [
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/sets:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "config_test.go",
        "keyring_test.go",
        "provider_test.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/credentialprovider",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/docker/api/types:go_default_library"],
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
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/aws:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/azure:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/gcp:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/credentialprovider/rancher:all-srcs",
    ],
    tags = ["automanaged"],
)
