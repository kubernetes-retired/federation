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
        "//vendor/github.com/pelletier/go-toml/cmd/tomljson:all-srcs",
        "//vendor/github.com/pelletier/go-toml/cmd/tomll:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
