load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["helpers.go"],
    cgo = True,
    importpath = "github.com/google/cadvisor/container/libcontainer",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/container:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["helpers_test.go"],
    importpath = "github.com/google/cadvisor/container/libcontainer",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/cgroups:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/system:go_default_library",
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
