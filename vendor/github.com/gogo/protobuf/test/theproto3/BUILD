load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = [
        "header.proto",
        "theproto3.proto",
    ],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/test/theproto3",
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
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/neither:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/theproto3/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
