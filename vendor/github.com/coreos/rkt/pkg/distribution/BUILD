load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "aciarchive.go",
        "appc.go",
        "cimd.go",
        "distribution.go",
        "docker.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/distribution",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/PuerkitoBio/purell:go_default_library",
        "//vendor/github.com/appc/docker2aci/lib/common:go_default_library",
        "//vendor/github.com/appc/spec/discovery:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common/labelsort:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "aciarchive_test.go",
        "appc_test.go",
        "docker_test.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/distribution",
    library = ":go_default_library",
    deps = ["//vendor/github.com/appc/spec/discovery:go_default_library"],
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
