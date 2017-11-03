load("@io_bazel_rules_go//go:def.bzl", "go_test")

go_test(
    name = "go_default_test",
    srcs = [
        "extension_test.go",
        "main_test.go",
    ],
    importpath = "github.com/golang/protobuf/protoc-gen-go/testdata",
    deps = [
        "//vendor/github.com/golang/protobuf/protoc-gen-go/testdata/multi:go_default_library",
        "//vendor/github.com/golang/protobuf/protoc-gen-go/testdata/my_test:go_default_library",
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
        "//vendor/github.com/golang/protobuf/protoc-gen-go/testdata/my_test:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
