load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["sysinfo.go"],
    importpath = "github.com/google/cadvisor/utils/sysinfo",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/google/cadvisor/utils/sysfs:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["sysinfo_test.go"],
    embed = [":go_default_library"],
    importpath = "github.com/google/cadvisor/utils/sysinfo",
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/google/cadvisor/utils/sysfs:go_default_library",
        "//vendor/github.com/google/cadvisor/utils/sysfs/fakesysfs:go_default_library",
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
