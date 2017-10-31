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
        "//vendor/k8s.io/kubernetes/test/e2e:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e_node:all-srcs",
        "//vendor/k8s.io/kubernetes/test/fixtures:all-srcs",
        "//vendor/k8s.io/kubernetes/test/images:all-srcs",
        "//vendor/k8s.io/kubernetes/test/integration:all-srcs",
        "//vendor/k8s.io/kubernetes/test/kubemark:all-srcs",
        "//vendor/k8s.io/kubernetes/test/list:all-srcs",
        "//vendor/k8s.io/kubernetes/test/soak/cauldron:all-srcs",
        "//vendor/k8s.io/kubernetes/test/soak/serve_hostnames:all-srcs",
        "//vendor/k8s.io/kubernetes/test/utils:all-srcs",
    ],
    tags = ["automanaged"],
)
