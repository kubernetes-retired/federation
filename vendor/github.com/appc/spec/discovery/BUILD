load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "discovery.go",
        "doc.go",
        "http.go",
        "parse.go",
    ],
    importpath = "github.com/appc/spec/discovery",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/common:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/golang.org/x/net/html:go_default_library",
        "//vendor/golang.org/x/net/html/atom:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "discovery_test.go",
        "http_test.go",
        "parse_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "github.com/appc/spec/discovery",
    library = ":go_default_library",
    deps = ["//vendor/github.com/appc/spec/schema/types:go_default_library"],
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
