load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/heketi/heketi",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/heketi/heketi/apps:go_default_library",
        "//vendor/github.com/heketi/heketi/apps/glusterfs:go_default_library",
        "//vendor/github.com/heketi/heketi/middleware:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/urfave/negroni:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
    ],
)

go_binary(
    name = "heketi",
    importpath = "github.com/heketi/heketi",
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
        "//vendor/github.com/heketi/heketi/apps:all-srcs",
        "//vendor/github.com/heketi/heketi/client/api/go-client:all-srcs",
        "//vendor/github.com/heketi/heketi/client/cli/go:all-srcs",
        "//vendor/github.com/heketi/heketi/executors:all-srcs",
        "//vendor/github.com/heketi/heketi/middleware:all-srcs",
        "//vendor/github.com/heketi/heketi/pkg/db:all-srcs",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:all-srcs",
        "//vendor/github.com/heketi/heketi/pkg/heketitest:all-srcs",
        "//vendor/github.com/heketi/heketi/pkg/kubernetes:all-srcs",
        "//vendor/github.com/heketi/heketi/pkg/utils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
