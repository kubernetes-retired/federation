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
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/password/allow:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/password/keystone:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/password/passwordfile:all-srcs",
    ],
    tags = ["automanaged"],
)
