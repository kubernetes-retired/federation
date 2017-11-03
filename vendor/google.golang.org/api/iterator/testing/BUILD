load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["testing.go"],
    importpath = "google.golang.org/api/iterator/testing",
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/api/iterator:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["testing_test.go"],
    importpath = "google.golang.org/api/iterator/testing",
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
