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
        "//vendor/github.com/docker/docker/integration-cli/checker:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/cli:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/daemon:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/environment:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/fixtures/load:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/fixtures/plugin:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/registry:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/request:all-srcs",
        "//vendor/github.com/docker/docker/integration-cli/requirement:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
