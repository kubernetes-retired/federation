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
        "//vendor/github.com/go-openapi/runtime/middleware/denco:all-srcs",
        "//vendor/github.com/go-openapi/runtime/middleware/header:all-srcs",
        "//vendor/github.com/go-openapi/runtime/middleware/untyped:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
