load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["oauth.go"],
    importpath = "go4.org/oauthutil",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/go4.org/wkfs:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
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
