load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["heketitest.go"],
    importpath = "github.com/heketi/heketi/pkg/heketitest",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/heketi/heketi/apps/glusterfs:go_default_library",
        "//vendor/github.com/heketi/heketi/middleware:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
        "//vendor/github.com/lpabon/godbc:go_default_library",
        "//vendor/github.com/urfave/negroni:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["heketitest_test.go"],
    importpath = "github.com/heketi/heketi/pkg/heketitest",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/heketi/heketi/client/api/go-client:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
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
