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
        "//vendor/github.com/json-iterator/go/output_tests/struct_tags/omitempty/ptr_string/string:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
