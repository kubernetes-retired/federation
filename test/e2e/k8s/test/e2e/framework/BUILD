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
        "//test/e2e/k8s/test/e2e/framework/ginkgowrapper:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
