load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "popcnt_generic.go",
        "sparse.go",
        "util.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "popcnt_amd64.go",
            "popcnt_amd64.s",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "popcnt_amd64.go",
            "popcnt_amd64.s",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "popcnt_amd64.go",
            "popcnt_amd64.s",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/tools/container/intsets",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["util_test.go"],
    importpath = "golang.org/x/tools/container/intsets",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["sparse_test.go"],
    importpath = "golang.org/x/tools/container/intsets_test",
    deps = [":go_default_library"],
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
