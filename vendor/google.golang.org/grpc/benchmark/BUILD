load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["benchmark.go"],
    importpath = "google.golang.org/grpc/benchmark",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/benchmark/grpc_testing:go_default_library",
        "//vendor/google.golang.org/grpc/grpclog:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["benchmark_test.go"],
    importpath = "google.golang.org/grpc/benchmark",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/benchmark/grpc_testing:go_default_library",
        "//vendor/google.golang.org/grpc/benchmark/stats:go_default_library",
        "//vendor/google.golang.org/grpc/grpclog:go_default_library",
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
        "//vendor/google.golang.org/grpc/benchmark/client:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/grpc_testing:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/server:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/stats:all-srcs",
        "//vendor/google.golang.org/grpc/benchmark/worker:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
