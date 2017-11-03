load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "backupdb.go",
        "namespace.go",
        "pv.go",
    ],
    importpath = "github.com/heketi/heketi/pkg/kubernetes",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/boltdb/bolt:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/glusterfs/api:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["backupdb_test.go"],
    importpath = "github.com/heketi/heketi/pkg/kubernetes",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/boltdb/bolt:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/client/clientset_generated/clientset/fake:go_default_library",
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
