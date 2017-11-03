load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["gcs.go"],
    importpath = "go4.org/wkfs/gcs",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go4.org/wkfs:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/golang.org/x/oauth2/google:go_default_library",
        "//vendor/google.golang.org/cloud:go_default_library",
        "//vendor/google.golang.org/cloud/compute/metadata:go_default_library",
        "//vendor/google.golang.org/cloud/storage:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["gcs_test.go"],
    importpath = "go4.org/wkfs/gcs",
    library = ":go_default_library",
    deps = [
        "//vendor/go4.org/wkfs:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/cloud/compute/metadata:go_default_library",
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
