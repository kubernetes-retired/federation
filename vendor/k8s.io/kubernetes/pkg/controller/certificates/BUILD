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
        "//vendor/k8s.io/kubernetes/pkg/controller/certificates/approver:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/certificates/cleaner:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/certificates/signer:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = [
        "//vendor/k8s.io/kubernetes/pkg/controller:__pkg__",
    ],
)
