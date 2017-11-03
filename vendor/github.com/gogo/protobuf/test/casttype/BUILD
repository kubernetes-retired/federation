load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["casttype.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["mytypes.go"],
    importpath = "github.com/gogo/protobuf/test/casttype",
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
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/neither:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/casttype/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
