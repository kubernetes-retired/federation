filegroup(
    name = "go_default_library_protos",
    srcs = ["plugin.proto"],
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
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-analyze:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-generator:all-srcs",
        "//vendor/github.com/googleapis/gnostic/plugins/gnostic-go-sample:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
