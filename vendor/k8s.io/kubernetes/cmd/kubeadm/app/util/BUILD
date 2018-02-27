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
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/apiclient:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/config:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/dryrun:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/pubkeypin:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/staticpod:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/util/token:all-srcs",
    ],
    tags = ["automanaged"],
)
