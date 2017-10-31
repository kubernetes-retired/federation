load("@io_bazel_rules_go//go:def.bzl", "go_library")

filegroup(
    name = "go_default_library_protos",
    srcs = ["flexvolume.proto"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "api_server.go",
        "client.go",
        "flexvolume.go",
        "flexvolume.pb.go",
        "flexvolume.pb.gw.go",
        "local_api_client.go",
    ],
    importpath = "github.com/libopenstorage/openstorage/pkg/flexvolume",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/protobuf/ptypes/empty:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/runtime:go_default_library",
        "//vendor/github.com/grpc-ecosystem/grpc-gateway/utilities:go_default_library",
        "//vendor/go.pedge.io/dlog:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/genproto/googleapis/api/annotations:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/grpclog:go_default_library",
        "//vendor/google.golang.org/grpc/status:go_default_library",
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
        "//vendor/github.com/libopenstorage/openstorage/pkg/flexvolume/cmd/flexvolume:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
