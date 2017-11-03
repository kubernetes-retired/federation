load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "os_default.go",
        "os_shared.go",
        "validator.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "os_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/appc/spec/ace",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
    ],
)

go_binary(
    name = "ace",
    importpath = "github.com/appc/spec/ace",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["validator_test.go"],
    importpath = "github.com/appc/spec/ace",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
