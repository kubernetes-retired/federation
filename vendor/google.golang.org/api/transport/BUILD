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
        "//vendor/google.golang.org/api/transport/bytestream:all-srcs",
        "//vendor/google.golang.org/api/transport/grpc:all-srcs",
        "//vendor/google.golang.org/api/transport/http:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
