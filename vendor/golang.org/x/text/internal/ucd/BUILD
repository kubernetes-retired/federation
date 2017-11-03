load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["ucd.go"],
    importpath = "golang.org/x/text/internal/ucd",
    visibility = ["//vendor/golang.org/x/text:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["ucd_test.go"],
    importpath = "golang.org/x/text/internal/ucd",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "golang.org/x/text/internal/ucd_test",
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
