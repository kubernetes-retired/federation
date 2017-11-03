load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/coreos/rkt/tests/rkt-monitor",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/shirou/gopsutil/load:go_default_library",
        "//vendor/github.com/shirou/gopsutil/process:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
    ],
)

go_binary(
    name = "rkt-monitor",
    importpath = "github.com/coreos/rkt/tests/rkt-monitor",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/cpu-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/log-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/mem-stresser:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor/sleeper:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
