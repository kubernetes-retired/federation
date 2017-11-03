load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "streamformatter.go",
        "streamwriter.go",
    ],
    importpath = "github.com/docker/docker/pkg/streamformatter",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/jsonmessage:go_default_library",
        "//vendor/github.com/docker/docker/pkg/progress:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "streamformatter_test.go",
        "streamwriter_test.go",
    ],
    importpath = "github.com/docker/docker/pkg/streamformatter",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/pkg/jsonmessage:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
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
