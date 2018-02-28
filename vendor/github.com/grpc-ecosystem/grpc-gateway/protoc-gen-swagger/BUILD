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
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/genswagger:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
