filegroup(
    name = "go_default_library_protos",
    srcs = [
        "a_bit_of_everything.proto",
        "echo_service.proto",
        "flow_combination.proto",
        "stream.proto",
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
