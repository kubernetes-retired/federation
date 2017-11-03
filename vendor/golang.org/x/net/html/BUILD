load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "const.go",
        "doc.go",
        "doctype.go",
        "entity.go",
        "escape.go",
        "foreign.go",
        "node.go",
        "parse.go",
        "render.go",
        "token.go",
    ],
    importpath = "golang.org/x/net/html",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/net/html/atom:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "entity_test.go",
        "escape_test.go",
        "node_test.go",
        "parse_test.go",
        "render_test.go",
        "token_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/net/html",
    library = ":go_default_library",
    deps = ["//vendor/golang.org/x/net/html/atom:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/net/html_test",
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
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/net/html/atom:all-srcs",
        "//vendor/golang.org/x/net/html/charset:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
