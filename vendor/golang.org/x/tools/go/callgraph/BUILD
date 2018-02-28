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
        "//vendor/golang.org/x/tools/go/callgraph/cha:all-srcs",
        "//vendor/golang.org/x/tools/go/callgraph/rta:all-srcs",
        "//vendor/golang.org/x/tools/go/callgraph/static:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
