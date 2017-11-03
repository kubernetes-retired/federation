load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["disco.go"],
    importpath = "google.golang.org/api/google-api-go-generator/internal/disco",
    visibility = ["//vendor/google.golang.org/api/google-api-go-generator:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["disco_test.go"],
    data = glob(["testdata/**"]),
    importpath = "google.golang.org/api/google-api-go-generator/internal/disco",
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
