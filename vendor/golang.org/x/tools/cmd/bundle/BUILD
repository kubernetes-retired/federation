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
        "//vendor/golang.org/x/tools/cmd/bundle/testdata/src/domain.name/importdecl:all-srcs",
        "//vendor/golang.org/x/tools/cmd/bundle/testdata/src/initial:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
