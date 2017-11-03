load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/test/oneof3",
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
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/neither:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/oneof3/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
