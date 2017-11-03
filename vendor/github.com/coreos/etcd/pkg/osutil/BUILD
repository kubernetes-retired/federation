load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "interrupt_unix.go",
        "osutil.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "interrupt_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/coreos/etcd/pkg/osutil",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/coreos/pkg/capnslog:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["osutil_test.go"],
    importpath = "github.com/coreos/etcd/pkg/osutil",
    library = ":go_default_library",
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
