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
        "//vendor/k8s.io/kubernetes/api/openapi-spec:all-srcs",
        "//vendor/k8s.io/kubernetes/api/swagger-spec:all-srcs",
    ],
    tags = ["automanaged"],
)
