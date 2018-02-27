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
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd/upgrade:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/cmd/util:all-srcs",
    ],
    tags = ["automanaged"],
)
