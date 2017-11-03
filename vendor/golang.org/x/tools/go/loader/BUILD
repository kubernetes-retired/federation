load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cgo.go",
        "cgo_pkgconfig.go",
        "doc.go",
        "loader.go",
        "util.go",
    ],
    importpath = "golang.org/x/tools/go/loader",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/ast/astutil:go_default_library",
        "//vendor/golang.org/x/tools/go/buildutil:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "example_test.go",
        "loader_test.go",
        "stdlib_test.go",
    ],
    importpath = "golang.org/x/tools/go/loader_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/tools/go/buildutil:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/tools/go/loader/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
