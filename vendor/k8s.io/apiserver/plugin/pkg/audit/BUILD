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
        "//vendor/k8s.io/apiserver/plugin/pkg/audit/log:all-srcs",
        "//vendor/k8s.io/apiserver/plugin/pkg/audit/webhook:all-srcs",
    ],
    tags = ["automanaged"],
)
