load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = [
        "df.proto",
        "dg.proto",
        "nc.proto",
        "ne.proto",
        "nx.proto",
    ],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/test/defaultconflict",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["nc_test.go"],
    importpath = "github.com/gogo/protobuf/test/defaultconflict",
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
