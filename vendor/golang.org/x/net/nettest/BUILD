load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "conntest.go",
        "conntest_go16.go",
        "conntest_go17.go",
    ],
    importpath = "golang.org/x/net/nettest",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["conntest_test.go"],
    importpath = "golang.org/x/net/nettest",
    library = ":go_default_library",
    deps = ["//vendor/golang.org/x/net/internal/nettest:go_default_library"],
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
