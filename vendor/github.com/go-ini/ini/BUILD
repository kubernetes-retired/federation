load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "ini.go",
        "struct.go",
    ],
    importpath = "github.com/go-ini/ini",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "ini_test.go",
        "struct_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "github.com/go-ini/ini",
    library = ":go_default_library",
    deps = ["//vendor/github.com/smartystreets/goconvey/convey:go_default_library"],
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
