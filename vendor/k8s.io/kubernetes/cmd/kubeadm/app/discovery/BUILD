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
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/discovery/file:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/discovery/https:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/discovery/token:all-srcs",
    ],
    tags = ["automanaged"],
)
