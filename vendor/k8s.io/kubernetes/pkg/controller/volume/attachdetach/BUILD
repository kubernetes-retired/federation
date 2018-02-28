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
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/cache:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/populator:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/reconciler:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/statusupdater:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/testing:all-srcs",
        "//vendor/k8s.io/kubernetes/pkg/controller/volume/attachdetach/util:all-srcs",
    ],
    tags = ["automanaged"],
)
