load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "devfs_nonlinux.go",
        "i2c.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "devfs.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/io/i2c",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/exp/io/i2c/driver:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["i2c_test.go"],
    importpath = "golang.org/x/exp/io/i2c",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "example_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/io/i2c_test",
    deps = select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            ":go_default_library",
        ],
        "//conditions:default": [],
    }),
)

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
        "//vendor/golang.org/x/exp/io/i2c/driver:all-srcs",
        "//vendor/golang.org/x/exp/io/i2c/example/displayip:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
