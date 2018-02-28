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
        "//vendor/k8s.io/apimachinery/pkg/apis/testapigroup/fuzzer:all-srcs",
        "//vendor/k8s.io/apimachinery/pkg/apis/testapigroup/install:all-srcs",
        "//vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1:all-srcs",
    ],
    tags = ["automanaged"],
)
