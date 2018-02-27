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
        "//vendor/k8s.io/apiserver/pkg/registry/generic:all-srcs",
        "//vendor/k8s.io/apiserver/pkg/registry/rest:all-srcs",
    ],
    tags = ["automanaged"],
)
