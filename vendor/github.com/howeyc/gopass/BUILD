load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "pass.go",
        "terminal.go",
    ],
    importpath = "github.com/howeyc/gopass",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/ssh/terminal:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["pass_test.go"],
    importpath = "github.com/howeyc/gopass",
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
