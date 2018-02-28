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
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakecontext:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakegit:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli/build/fakestorage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
