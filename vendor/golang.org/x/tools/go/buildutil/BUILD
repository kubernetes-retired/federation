load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "allpackages.go",
        "fakecontext.go",
        "overlay.go",
        "tags.go",
        "util.go",
    ],
    importpath = "golang.org/x/tools/go/buildutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "allpackages_test.go",
        "overlay_test.go",
        "tags_test.go",
        "util_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "util_windows_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/tools/go/buildutil_test",
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
