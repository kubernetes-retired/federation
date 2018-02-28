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
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/password:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/request/basicauth:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/token/oidc:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/token/tokentest:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/authenticator/token/webhook:all-srcs",
    ],
    tags = ["automanaged"],
)
