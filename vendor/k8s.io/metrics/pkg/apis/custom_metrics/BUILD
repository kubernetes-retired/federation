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
        "//vendor/k8s.io/metrics/pkg/apis/custom_metrics/install:all-srcs",
        "//vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1:all-srcs",
    ],
    tags = ["automanaged"],
)
