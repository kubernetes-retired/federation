load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "golang.org/x/tools/cmd/guru/testdata/src/referrers",
    visibility = ["//visibility:private"],
    deps = ["//vendor/lib:go_default_library"],
)

go_binary(
    name = "referrers",
    importpath = "golang.org/x/tools/cmd/guru/testdata/src/referrers",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["int_test.go"],
    importpath = "golang.org/x/tools/cmd/guru/testdata/src/referrers",
    library = ":go_default_library",
    deps = ["//vendor/lib:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["ext_test.go"],
    importpath = "golang.org/x/tools/cmd/guru/testdata/src/referrers_test",
    deps = [
        "//vendor/lib:go_default_library",
        "//vendor/referrers:go_default_library",
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
