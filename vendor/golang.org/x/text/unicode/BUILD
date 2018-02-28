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
        "//vendor/golang.org/x/text/unicode/bidi:all-srcs",
        "//vendor/golang.org/x/text/unicode/cldr:all-srcs",
        "//vendor/golang.org/x/text/unicode/norm:all-srcs",
        "//vendor/golang.org/x/text/unicode/rangetable:all-srcs",
        "//vendor/golang.org/x/text/unicode/runenames:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
