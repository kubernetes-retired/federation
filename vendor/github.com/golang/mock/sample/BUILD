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
        "//vendor/github.com/golang/mock/sample/imp1:all-srcs",
        "//vendor/github.com/golang/mock/sample/imp2:all-srcs",
        "//vendor/github.com/golang/mock/sample/imp3:all-srcs",
        "//vendor/github.com/golang/mock/sample/imp4:all-srcs",
        "//vendor/github.com/golang/mock/sample/mock_user:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
