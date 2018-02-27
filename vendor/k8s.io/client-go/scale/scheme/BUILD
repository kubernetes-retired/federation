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
        "//vendor/k8s.io/client-go/scale/scheme/autoscalingv1:all-srcs",
        "//vendor/k8s.io/client-go/scale/scheme/extensionsint:all-srcs",
        "//vendor/k8s.io/client-go/scale/scheme/extensionsv1beta1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
