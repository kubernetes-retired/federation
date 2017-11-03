load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/vanity/test",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["vanity_test.go"],
    importpath = "github.com/gogo/protobuf/vanity/test",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/gogo/protobuf/vanity/test/fast:go_default_library",
        "//vendor/github.com/gogo/protobuf/vanity/test/faster:go_default_library",
        "//vendor/github.com/gogo/protobuf/vanity/test/slick:go_default_library",
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
        "//vendor/github.com/gogo/protobuf/vanity/test/fast:all-srcs",
        "//vendor/github.com/gogo/protobuf/vanity/test/faster:all-srcs",
        "//vendor/github.com/gogo/protobuf/vanity/test/slick:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
