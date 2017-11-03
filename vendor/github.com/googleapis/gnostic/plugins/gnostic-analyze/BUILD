load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-analyze",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/googleapis/gnostic/OpenAPIv2:go_default_library",
        "//vendor/github.com/googleapis/gnostic/plugins:go_default_library",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze/statistics:go_default_library",
    ],
)

go_binary(
    name = "gnostic-analyze",
    importpath = "github.com/googleapis/gnostic/plugins/gnostic-analyze",
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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze/statistics:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze/summarize:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
