load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "noop.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "tar.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/Microsoft/go-winio/backuptar",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Microsoft/go-winio:go_default_library",
            "//vendor/github.com/Microsoft/go-winio/archive/tar:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = ["tar_test.go"],
    importpath = "github.com/Microsoft/go-winio/backuptar",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/Microsoft/go-winio:go_default_library",
        "//vendor/github.com/Microsoft/go-winio/archive/tar:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
