load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "aws.go",
        "aws_storage.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/aws",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/session:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/ec2:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/opsworks:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/chaos:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/proto/time:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/common:go_default_library",
        "//vendor/github.com/portworx/kvdb:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["aws_test.go"],
    importpath = "github.com/libopenstorage/openstorage/volume/drivers/aws",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/api:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/volume/drivers/test:go_default_library",
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
