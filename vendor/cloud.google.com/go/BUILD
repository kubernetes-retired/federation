load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["cloud.go"],
    importpath = "cloud.google.com/go",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["license_test.go"],
    importpath = "cloud.google.com/go",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["authexample_test.go"],
    importpath = "cloud.google.com/go_test",
    deps = [
        "//vendor/cloud.google.com/go/datastore:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/api/option:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/cloud.google.com/go/bigquery:all-srcs",
        "//vendor/cloud.google.com/go/bigtable:all-srcs",
        "//vendor/cloud.google.com/go/compute/metadata:all-srcs",
        "//vendor/cloud.google.com/go/container:all-srcs",
        "//vendor/cloud.google.com/go/datastore:all-srcs",
        "//vendor/cloud.google.com/go/errors:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigquery/concat_table:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigquery/load:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigquery/query:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigquery/read:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigtable/helloworld:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigtable/search:all-srcs",
        "//vendor/cloud.google.com/go/examples/bigtable/usercounter:all-srcs",
        "//vendor/cloud.google.com/go/examples/storage/appengine:all-srcs",
        "//vendor/cloud.google.com/go/examples/storage/appenginevm:all-srcs",
        "//vendor/cloud.google.com/go/internal:all-srcs",
        "//vendor/cloud.google.com/go/language/apiv1beta1:all-srcs",
        "//vendor/cloud.google.com/go/logging:all-srcs",
        "//vendor/cloud.google.com/go/preview/logging:all-srcs",
        "//vendor/cloud.google.com/go/pubsub:all-srcs",
        "//vendor/cloud.google.com/go/speech/apiv1beta1:all-srcs",
        "//vendor/cloud.google.com/go/storage:all-srcs",
        "//vendor/cloud.google.com/go/trace:all-srcs",
        "//vendor/cloud.google.com/go/vision:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
