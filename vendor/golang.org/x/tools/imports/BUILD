load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "fastwalk.go",
        "fastwalk_portable.go",
        "fix.go",
        "imports.go",
        "sortimports.go",
        "zstdlib.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "fastwalk_dirent_ino.go",
            "fastwalk_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "fastwalk_dirent_ino.go",
            "fastwalk_unix.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/tools/imports",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/go/ast/astutil:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "fastwalk_test.go",
        "fix_test.go",
    ],
    importpath = "golang.org/x/tools/imports",
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
