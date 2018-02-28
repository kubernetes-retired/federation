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
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/examples/server/cmd/example-server:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
