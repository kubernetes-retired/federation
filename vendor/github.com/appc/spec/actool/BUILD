load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "actool.go",
        "build.go",
        "discover.go",
        "doc.go",
        "help.go",
        "manifest.go",
        "validate.go",
        "version.go",
    ],
    importpath = "github.com/appc/spec/actool",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/discovery:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
    ],
)

go_binary(
    name = "actool",
    importpath = "github.com/appc/spec/actool",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
