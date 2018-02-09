load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "discovery.go",
        "env.go",
        "http.go",
        "vcs.go",
    ],
    importpath = "golang.org/x/tools/go/vcs",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["vcs_test.go"],
    embed = [":go_default_library"],
    importpath = "golang.org/x/tools/go/vcs",
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
