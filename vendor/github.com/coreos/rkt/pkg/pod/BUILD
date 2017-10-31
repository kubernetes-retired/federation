load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "pods.go",
        "sandbox.go",
        "uuid.go",
        "wait.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/pod",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/networking/netinfo:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/label:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/sys:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/github.com/pborman/uuid:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "pods_test.go",
        "wait_test.go",
    ],
    importpath = "github.com/coreos/rkt/pkg/pod",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
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
