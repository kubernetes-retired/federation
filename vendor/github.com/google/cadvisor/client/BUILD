load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["client.go"],
    importpath = "github.com/google/cadvisor/client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["client_test.go"],
    importpath = "github.com/google/cadvisor/client",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/google/cadvisor/info/v1:go_default_library",
        "//vendor/github.com/google/cadvisor/info/v1/test:go_default_library",
        "//vendor/github.com/kr/pretty:go_default_library",
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
        "//vendor/github.com/google/cadvisor/client/clientexample:all-srcs",
        "//vendor/github.com/google/cadvisor/client/v2:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
