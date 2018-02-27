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
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor:all-srcs",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/generator:all-srcs",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/grpc:all-srcs",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin:all-srcs",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
