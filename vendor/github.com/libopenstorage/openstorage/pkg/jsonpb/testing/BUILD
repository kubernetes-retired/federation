load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

filegroup(
    name = "go_default_library_protos",
    srcs = ["testing.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "testing.go",
        "testing.pb.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/pkg/jsonpb/testing",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/timestamp:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/jsonpb:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["testing_test.go"],
    importpath = "github.com/libopenstorage/openstorage/pkg/jsonpb/testing",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/golang/protobuf/ptypes/timestamp:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/jsonpb:go_default_library",
        "//vendor/github.com/libopenstorage/openstorage/pkg/proto/time:go_default_library",
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
