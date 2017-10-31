load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "line_parsers.go",
        "parser.go",
        "split_command.go",
    ],
    importpath = "github.com/docker/docker/builder/dockerfile/parser",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/builder/dockerfile/command:go_default_library",
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "json_test.go",
        "line_parsers_test.go",
        "parser_test.go",
    ],
    importpath = "github.com/docker/docker/builder/dockerfile/parser",
    library = ":go_default_library",
    deps = [
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/docker/docker/builder/dockerfile/parser/dumper:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
