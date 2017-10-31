load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["one.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/test/oneof",
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
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/neither:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
