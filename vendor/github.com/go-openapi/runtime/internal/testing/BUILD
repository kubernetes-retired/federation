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
        "//vendor/github.com/go-openapi/runtime/internal/testing/petstore:all-srcs",
        "//vendor/github.com/go-openapi/runtime/internal/testing/simplepetstore:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
