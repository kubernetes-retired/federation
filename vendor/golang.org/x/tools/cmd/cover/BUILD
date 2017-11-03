load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cover.go",
        "doc.go",
        "func.go",
        "html.go",
    ],
    importpath = "golang.org/x/tools/cmd/cover",
    visibility = ["//visibility:private"],
    deps = ["//vendor/golang.org/x/tools/cover:go_default_library"],
)

go_binary(
    name = "cover",
    importpath = "golang.org/x/tools/cmd/cover",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["cover_test.go"],
    importpath = "golang.org/x/tools/cmd/cover_test",
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/tools/cmd/cover/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
