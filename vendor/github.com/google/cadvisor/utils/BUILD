load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "path.go",
        "timed_store.go",
        "utils.go",
    ],
    importpath = "github.com/google/cadvisor/utils",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["timed_store_test.go"],
    importpath = "github.com/google/cadvisor/utils",
    library = ":go_default_library",
    deps = ["//vendor/github.com/stretchr/testify/assert:go_default_library"],
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
        "//vendor/github.com/google/cadvisor/utils/cloudinfo:all-srcs",
        "//vendor/github.com/google/cadvisor/utils/cpuload:all-srcs",
        "//vendor/github.com/google/cadvisor/utils/docker:all-srcs",
        "//vendor/github.com/google/cadvisor/utils/oomparser:all-srcs",
        "//vendor/github.com/google/cadvisor/utils/sysfs:all-srcs",
        "//vendor/github.com/google/cadvisor/utils/sysinfo:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
