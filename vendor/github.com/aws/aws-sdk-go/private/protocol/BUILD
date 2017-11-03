load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "idempotency.go",
        "unmarshal.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/private/protocol",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/aws/aws-sdk-go/aws/request:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "idempotency_test.go",
        "protocol_test.go",
        "unmarshal_test.go",
    ],
    importpath = "github.com/aws/aws-sdk-go/private/protocol_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/client/metadata:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/request:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/awstesting:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/ec2query:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/jsonrpc:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/query:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/rest:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/restjson:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/restxml:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/ec2query:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/json/jsonutil:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/jsonrpc:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/query:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/rest:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/restjson:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/restxml:all-srcs",
        "//vendor/github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
