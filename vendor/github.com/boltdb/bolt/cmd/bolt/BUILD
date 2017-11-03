load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/boltdb/bolt/cmd/bolt",
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/boltdb/bolt:go_default_library"],
)

go_binary(
    name = "bolt",
    importpath = "github.com/boltdb/bolt/cmd/bolt",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["main_test.go"],
    importpath = "github.com/boltdb/bolt/cmd/bolt_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/boltdb/bolt:go_default_library",
    ],
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
