filegroup(
    name = "go_default_library_protos",
    srcs = ["castvalue.proto"],
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
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/castvalue/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
