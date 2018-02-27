filegroup(
    name = "go_default_library_protos",
    srcs = ["OpenAPIv3.proto"],
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
        "//vendor/github.com/googleapis/gnostic/OpenAPIv3/schema-generator:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
