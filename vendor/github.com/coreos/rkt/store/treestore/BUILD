load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["tree.go"],
    importpath = "github.com/coreos/rkt/store/treestore",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/pkg/acirenderer:go_default_library",
        "//vendor/github.com/appc/spec/pkg/tarheader:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/aci:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/fileutil:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/coreos/rkt/store/imagestore:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["tree_test.go"],
    importpath = "github.com/coreos/rkt/store/treestore",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/aci:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/aci/acitest:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/coreos/rkt/store/imagestore:go_default_library",
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
