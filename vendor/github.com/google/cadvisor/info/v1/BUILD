load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "container.go",
        "docker.go",
        "machine.go",
        "metric.go",
    ],
    importpath = "github.com/google/cadvisor/info/v1",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["container_test.go"],
    importpath = "github.com/google/cadvisor/info/v1",
    library = ":go_default_library",
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
        "//vendor/github.com/google/cadvisor/info/v1/test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
