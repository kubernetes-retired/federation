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
        "//vendor/google.golang.org/grpc/benchmark/client:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/grpc_testing:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/server:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/stats:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/worker:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
