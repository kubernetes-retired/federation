filegroup(
    name = "go_default_library_protos",
    srcs = [
        "eb.proto",
        "ec.proto",
        "ee.proto",
        "em.proto",
        "en.proto",
        "er.proto",
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
