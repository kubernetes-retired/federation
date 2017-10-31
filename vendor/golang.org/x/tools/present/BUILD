load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "args.go",
        "background.go",
        "caption.go",
        "code.go",
        "doc.go",
        "html.go",
        "iframe.go",
        "image.go",
        "link.go",
        "parse.go",
        "style.go",
        "video.go",
    ],
    importpath = "golang.org/x/tools/present",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "code_test.go",
        "link_test.go",
        "style_test.go",
    ],
    importpath = "golang.org/x/tools/present",
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
