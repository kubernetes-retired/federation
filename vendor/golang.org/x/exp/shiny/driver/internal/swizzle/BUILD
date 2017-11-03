load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "swizzle_common.go",
        "swizzle_other.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "swizzle_amd64.go",
            "swizzle_amd64.s",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "swizzle_amd64.go",
            "swizzle_amd64.s",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "swizzle_amd64.go",
            "swizzle_amd64.s",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/exp/shiny/driver/internal/swizzle",
    visibility = ["//vendor/golang.org/x/exp/shiny/driver:__subpackages__"],
)

go_test(
    name = "go_default_test",
    srcs = ["swizzle_test.go"],
    importpath = "golang.org/x/exp/shiny/driver/internal/swizzle",
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
