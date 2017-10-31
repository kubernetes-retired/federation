load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ABitOfEverythingNested.go",
        "ABitOfEverythingServiceApi.go",
        "ExamplepbABitOfEverything.go",
        "ExamplepbNumericEnum.go",
        "NestedDeepEnum.go",
        "ProtobufEmpty.go",
        "Sub2IdMessage.go",
        "SubStringMessage.go",
    ],
    importpath = "github.com/grpc-ecosystem/grpc-gateway/examples/clients/abe",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/dghubble/sling:go_default_library"],
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
