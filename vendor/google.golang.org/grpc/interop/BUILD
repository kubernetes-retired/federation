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
        "//vendor/google.golang.org/grpc/interop/client:all-srcs",
        "//vendor/google.golang.org/grpc/interop/grpc_testing:all-srcs",
        "//vendor/google.golang.org/grpc/interop/http2:all-srcs",
        "//vendor/google.golang.org/grpc/interop/server:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
