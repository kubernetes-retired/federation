package(default_visibility = ["//visibility:public"])

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
        "//vendor/k8s.io/kubernetes/test/e2e/chaosmonkey:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/common:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/framework:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/generated:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/manifest:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/perftype:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/testing-manifests:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/builder:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node/services:all-srcs",
        "//vendor/k8s.io/kubernetes/test/fixtures:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration/framework:all-srcs",
        "//vendor/k8s.io/kubernetes/test/utils:all-srcs",
    ],
    tags = ["automanaged"],
)
