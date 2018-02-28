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
        "//vendor/golang.org/x/text/secure/bidirule:all-srcs",
        "//vendor/golang.org/x/text/secure/precis:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
