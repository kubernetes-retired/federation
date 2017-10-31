load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["fmtcmd.go"],
    importpath = "github.com/hashicorp/hcl/hcl/fmtcmd",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/hcl/hcl/printer:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["fmtcmd_test.go"],
    importpath = "github.com/hashicorp/hcl/hcl/fmtcmd",
    library = ":go_default_library",
    deps = ["//vendor/github.com/hashicorp/hcl/testhelper:go_default_library"],
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
