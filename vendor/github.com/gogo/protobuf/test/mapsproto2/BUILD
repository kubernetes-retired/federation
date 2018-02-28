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
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/both:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/marshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/neither:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/unmarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/unsafeboth:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/unsafemarshaler:all-srcs",
        "//vendor/github.com/gogo/protobuf/test/mapsproto2/combos/unsafeunmarshaler:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
