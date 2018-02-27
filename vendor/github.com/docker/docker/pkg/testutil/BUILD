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
        "//vendor/github.com/docker/docker/pkg/testutil/cmd:all-srcs",
        "//vendor/github.com/docker/docker/pkg/testutil/golden:all-srcs",
        "//vendor/github.com/docker/docker/pkg/testutil/tempfile:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
