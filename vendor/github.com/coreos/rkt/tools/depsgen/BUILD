load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "globcmd.go",
        "gocmd.go",
        "kvcmd.go",
        "main.go",
        "mkfile.go",
        "util.go",
    ],
    importpath = "github.com/coreos/rkt/tools/depsgen",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/coreos/rkt/tools/common:go_default_library",
        "//vendor/github.com/coreos/rkt/tools/common/filelist:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_binary(
    name = "depsgen",
    importpath = "github.com/coreos/rkt/tools/depsgen",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
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
