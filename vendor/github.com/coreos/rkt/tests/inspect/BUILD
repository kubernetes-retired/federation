load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["inspect.go"],
    importpath = "github.com/coreos/rkt/tests/inspect",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/pkg/device:go_default_library",
        "//vendor/github.com/coreos/rkt/common/cgroup:go_default_library",
        "//vendor/github.com/coreos/rkt/common/cgroup/v1:go_default_library",
        "//vendor/github.com/coreos/rkt/common/cgroup/v2:go_default_library",
        "//vendor/github.com/coreos/rkt/tests/testutils:go_default_library",
        "//vendor/github.com/syndtr/gocapability/capability:go_default_library",
        "//vendor/golang.org/x/sys/unix:go_default_library",
    ],
)

go_binary(
    name = "inspect",
    importpath = "github.com/coreos/rkt/tests/inspect",
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
