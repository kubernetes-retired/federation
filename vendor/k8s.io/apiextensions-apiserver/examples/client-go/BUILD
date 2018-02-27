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
        "//vendor/k8s.io/apiextensions-apiserver/examples/client-go/apis/cr/v1:all-srcs",
        "//vendor/k8s.io/apiextensions-apiserver/examples/client-go/client:all-srcs",
        "//vendor/k8s.io/apiextensions-apiserver/examples/client-go/controller:all-srcs",
    ],
    tags = ["automanaged"],
)
