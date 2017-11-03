load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["db.go"],
    importpath = "github.com/coreos/rkt/store/db",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/github.com/cznic/ql/driver:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["db_test.go"],
    importpath = "github.com/coreos/rkt/store/db",
    library = ":go_default_library",
    deps = ["//vendor/github.com/coreos/rkt/tests/testutils:go_default_library"],
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
