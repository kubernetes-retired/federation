filegroup(
    name = "go_default_library_protos",
    srcs = ["indeximport.proto"],
    visibility = ["//visibility:public"],
)

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
        "//vendor/github.com/gogo/protobuf/test/indeximport-issue72/index:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
