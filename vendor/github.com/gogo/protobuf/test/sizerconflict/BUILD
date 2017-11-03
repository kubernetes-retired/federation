load("@io_bazel_rules_go//go:def.bzl", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = ["sizerconflict.proto"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["sizerconflict_test.go"],
    importpath = "github.com/gogo/protobuf/test/sizerconflict",
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
