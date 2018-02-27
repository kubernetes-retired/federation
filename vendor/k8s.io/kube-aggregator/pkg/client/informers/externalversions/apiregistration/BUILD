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
        "//vendor/k8s.io/kube-aggregator/pkg/client/informers/externalversions/apiregistration/v1beta1:all-srcs",
    ],
    tags = ["automanaged"],
)
