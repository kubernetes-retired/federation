filegroup(
    name = "go_default_library_protos",
    srcs = ["tarsum.proto"],
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
        "//vendor/github.com/docker/docker/builder/remotecontext/git:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
