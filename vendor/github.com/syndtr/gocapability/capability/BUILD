load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "capability.go",
        "capability_noop.go",
        "enum.go",
        "enum_gen.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "capability_linux.go",
            "syscall_linux.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/syndtr/gocapability/capability",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["capability_test.go"],
    importpath = "github.com/syndtr/gocapability/capability",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/syndtr/gocapability/capability/enumgen:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
