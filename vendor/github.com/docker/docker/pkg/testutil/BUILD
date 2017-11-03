load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "helpers.go",
        "pkg.go",
        "utils.go",
    ],
    importpath = "github.com/docker/docker/pkg/testutil",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/stringutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["utils_test.go"],
    importpath = "github.com/docker/docker/pkg/testutil",
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
        "//vendor/github.com/docker/docker/pkg/testutil/cmd:all-srcs",
        "//vendor/github.com/docker/docker/pkg/testutil/golden:all-srcs",
        "//vendor/github.com/docker/docker/pkg/testutil/tempfile:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
