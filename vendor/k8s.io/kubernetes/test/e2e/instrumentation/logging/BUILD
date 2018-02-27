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
        "//vendor/k8s.io/kubernetes/test/e2e/instrumentation/logging/elasticsearch:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/instrumentation/logging/stackdrvier:all-srcs",
        "//vendor/k8s.io/kubernetes/test/e2e/instrumentation/logging/utils:all-srcs",
    ],
    tags = ["automanaged"],
)
