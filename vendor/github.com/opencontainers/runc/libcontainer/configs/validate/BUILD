load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "rootless.go",
        "validator.go",
    ],
    importpath = "github.com/opencontainers/runc/libcontainer/configs/validate",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
        "//vendor/github.com/opencontainers/selinux/go-selinux:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["rootless_test.go"],
    importpath = "github.com/opencontainers/runc/libcontainer/configs/validate",
    library = ":go_default_library",
    deps = ["//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = ["validator_test.go"],
    importpath = "github.com/opencontainers/runc/libcontainer/configs/validate_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/configs:go_default_library",
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
