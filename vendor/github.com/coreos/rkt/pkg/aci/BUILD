load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "aci.go",
        "render.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/aci",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/pkg/acirenderer:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/tar:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/user:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp:go_default_library",
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
        "//vendor/github.com/coreos/rkt/pkg/aci/acitest:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
