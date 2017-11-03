load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "emptyvfs.go",
        "namespace.go",
        "os.go",
        "vfs.go",
    ],
    importpath = "golang.org/x/tools/godoc/vfs",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["emptyvfs_test.go"],
    importpath = "golang.org/x/tools/godoc/vfs_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/tools/godoc/vfs/mapfs:go_default_library",
    ],
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
        "//vendor/golang.org/x/tools/godoc/vfs/gatefs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/httpfs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/mapfs:all-srcs",
        "//vendor/golang.org/x/tools/godoc/vfs/zipfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
