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
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor:all-srcs",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/generator:all-srcs",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/grpc:all-srcs",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/plugin:all-srcs",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
