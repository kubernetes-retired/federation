load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "authorization_config.pb.go",
        "experimental.pb.go",
    ],
    importpath = "google.golang.org/genproto/googleapis/api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
        "//vendor/google.golang.org/genproto/googleapis/api/annotations:go_default_library",
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
        "//vendor/google.golang.org/genproto/googleapis/api/annotations:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/configchange:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/distribution:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/httpbody:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/label:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/metric:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/monitoredres:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/serviceconfig:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/servicecontrol/v1:all-srcs",
        "//vendor/google.golang.org/genproto/googleapis/api/servicemanagement/v1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
