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
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/testdata/my_test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
