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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/encode-templates:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/apis_guru:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/bookstore:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator/examples/v2.0/xkcd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
