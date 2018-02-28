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
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/fruit.io/banana:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/fruit.io/orange:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/fruit.io/pear:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/new.com/one:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/old.com/one:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/titanic.biz/bar:all-srcs",
        "//vendor/golang.org/x/tools/cmd/fiximports/testdata/src/titanic.biz/foo:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
