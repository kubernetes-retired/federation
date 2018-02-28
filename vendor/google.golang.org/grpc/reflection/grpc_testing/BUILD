filegroup(
    name = "go_default_library_protos",
    srcs = [
        "proto2.proto",
        "proto2_ext.proto",
        "proto2_ext2.proto",
        "test.proto",
    ],
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
