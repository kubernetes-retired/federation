filegroup(
    name = "go_default_library_protos",
    srcs = ["flexvolume.proto"],
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
        "//vendor/github.com/libopenstorage/openstorage/pkg/flexvolume/cmd/flexvolume:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
