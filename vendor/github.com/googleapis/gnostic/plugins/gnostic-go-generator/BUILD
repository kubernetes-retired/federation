load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "funcmap.go",
        "gofmt.go",
        "main.go",
        "renderer.go",
        "templates.go",
    ],
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/googleapis/gnostic/OpenAPIv2:go_default_library",
        "//vendor/github.com/googleapis/gnostic/plugins:go_default_library",
    ],
)

go_binary(
    name = "gnostic-go-generator",
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-go-generator",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/encode-templates:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/apis_guru:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/xkcd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
