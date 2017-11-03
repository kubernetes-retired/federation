load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "image.go",
        "kind.go",
        "pod.go",
        "version.go",
    ],
    importpath = "github.com/appc/spec/schema",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/go4.org/errorutil:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "image_test.go",
        "pod_test.go",
    ],
    importpath = "github.com/appc/spec/schema",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/appc/spec/schema/common:all-srcs",
        "//vendor/github.com/appc/spec/schema/lastditch:all-srcs",
        "//vendor/github.com/appc/spec/schema/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
