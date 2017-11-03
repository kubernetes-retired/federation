load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "resolver.go",
        "uid_range.go",
        "util.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/user",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/group:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/passwd:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["resolver_test.go"],
    importpath = "github.com/coreos/rkt/pkg/user_test",
    deps = [":go_default_library"],
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
