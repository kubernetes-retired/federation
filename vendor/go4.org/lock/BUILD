load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "lock.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "lock_darwin_amd64.go",
            "lock_sigzero.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "lock_linux_amd64.go",
            "lock_sigzero.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "go4.org/lock",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["lock_test.go"],
    importpath = "go4.org/lock",
    library = ":go_default_library",
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
