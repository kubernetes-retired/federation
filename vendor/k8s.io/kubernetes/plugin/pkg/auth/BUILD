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
        "//vendor/k8s.io/kubernetes/plugin/pkg/auth/authenticator/token/bootstrap:all-srcs",
        "//vendor/k8s.io/kubernetes/plugin/pkg/auth/authorizer:all-srcs",
    ],
    tags = ["automanaged"],
)
