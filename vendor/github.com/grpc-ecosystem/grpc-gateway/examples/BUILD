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
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/clients/abe:all-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/clients/echo:all-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/examplepb:all-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/server:all-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/sub:all-srcs",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/sub2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
