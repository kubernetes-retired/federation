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
        "//vendor/github.com/onsi/gomega/gstruct/errors:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
