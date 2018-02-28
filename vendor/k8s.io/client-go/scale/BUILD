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
        "//vendor/k8s.io/client-go/scale/fake:all-srcs",
        "//vendor/k8s.io/client-go/scale/scheme:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
