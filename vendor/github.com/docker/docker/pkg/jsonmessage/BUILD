load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["jsonmessage.go"],
    importpath = "github.com/docker/docker/pkg/jsonmessage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Nvveen/Gotty:go_default_library",
        "//vendor/github.com/docker/docker/pkg/jsonlog:go_default_library",
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["jsonmessage_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/docker/docker/pkg/jsonmessage",
    deps = [
        "//vendor/github.com/docker/docker/pkg/jsonlog:go_default_library",
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
    ],
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
