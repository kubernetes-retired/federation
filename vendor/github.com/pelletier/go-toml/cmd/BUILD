load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["test_program.go"],
    importpath = "github.com/pelletier/go-toml/cmd",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/pelletier/go-toml:go_default_library"],
)

go_binary(
    name = "cmd",
    importpath = "github.com/pelletier/go-toml/cmd",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
        "//vendor/github.com/pelletier/go-toml/cmd/tomljson:all-srcs",
        "//vendor/github.com/pelletier/go-toml/cmd/tomll:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
