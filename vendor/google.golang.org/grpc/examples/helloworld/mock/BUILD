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
        "//vendor/google.golang.org/grpc/examples/helloworld/mock/mock_helloworld:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
