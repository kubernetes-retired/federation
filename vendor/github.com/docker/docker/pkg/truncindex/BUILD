load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["truncindex.go"],
    importpath = "github.com/docker/docker/pkg/truncindex",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/tchap/go-patricia/patricia:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["truncindex_test.go"],
    importpath = "github.com/docker/docker/pkg/truncindex",
    library = ":go_default_library",
    deps = ["//vendor/github.com/docker/docker/pkg/stringid:go_default_library"],
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
