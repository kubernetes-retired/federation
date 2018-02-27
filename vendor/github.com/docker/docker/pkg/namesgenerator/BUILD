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
        "//vendor/github.com/docker/docker/pkg/namesgenerator/cmd/names-generator:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
