load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["syncdebug.go"],
    importpath = "go4.org/syncutil/syncdebug",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go4.org/strutil:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["syncdebug_test.go"],
    importpath = "go4.org/syncutil/syncdebug",
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
