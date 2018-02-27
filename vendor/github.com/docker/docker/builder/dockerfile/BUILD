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
        "//vendor/github.com/docker/docker/builder/dockerfile/command:all-srcs",
        "//vendor/github.com/docker/docker/builder/dockerfile/parser:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
