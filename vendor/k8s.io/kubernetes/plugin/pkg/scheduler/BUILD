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
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/algorithm:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/algorithmprovider:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/api:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/core:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/factory:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/metrics:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/testing:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/scheduler/util:all-srcs",
    ],
    tags = ["automanaged"],
)
