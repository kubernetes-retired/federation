load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["cloudlaunch.go"],
    importpath = "go4.org/cloud/cloudlaunch",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go4.org/cloud/google/gceutil:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/golang.org/x/oauth2/google:go_default_library",
        "//vendor/google.golang.org/api/compute/v1:go_default_library",
        "//vendor/google.golang.org/api/googleapi:go_default_library",
        "//vendor/google.golang.org/api/storage/v1:go_default_library",
        "//vendor/google.golang.org/cloud:go_default_library",
        "//vendor/google.golang.org/cloud/storage:go_default_library",
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
