load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "alert.go",
        "alert_kvdb.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/alert",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/proto/time:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["alert_kvdb_test.go"],
    importpath = "github.com/libopenstorage/openstorage/alert",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/proto/time:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/github.com/portworx/kvdb/mem:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
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
