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
        "//vendor/k8s.io/kubernetes/cmd/gendocs:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/genkubedocs:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/genman:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/genswaggertypedocs:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/genutils:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/genyaml:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kube-apiserver:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kube-proxy:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm:all-srcs",
        "//vendor/k8s.io/kubernetes/cmd/kubelet:all-srcs",
    ],
    tags = ["automanaged"],
)
