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
        "//vendor/golang.org/x/tools/go/ssa/interp:all-srcs",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
