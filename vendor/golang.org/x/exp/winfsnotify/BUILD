load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "winfsnotify.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/winfsnotify",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "winfsnotify_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/winfsnotify",
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
