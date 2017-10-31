load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "osutil.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "exec_sysctl.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "exec_procfs.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "exec_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "go4.org/osutil",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["exec_test.go"],
    importpath = "go4.org/osutil",
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
