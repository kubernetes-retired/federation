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
        "//vendor/google.golang.org/api/google-api-go-generator/internal/disco:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
