load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["generator.go"],
    importpath = "github.com/golang/protobuf/protoc-gen-go/generator",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor:go_default_library",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/plugin:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["name_test.go"],
    importpath = "github.com/golang/protobuf/protoc-gen-go/generator",
    library = ":go_default_library",
    deps = ["//vendor/github.com/golang/protobuf/protoc-gen-go/descriptor:go_default_library"],
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
