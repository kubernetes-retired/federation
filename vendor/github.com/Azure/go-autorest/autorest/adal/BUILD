load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "devicetoken.go",
        "persist.go",
        "sender.go",
        "token.go",
    ],
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/dgrijalva/jwt-go:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "config_test.go",
        "devicetoken_test.go",
        "persist_test.go",
        "token_test.go",
    ],
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    library = ":go_default_library",
    deps = ["//vendor/github.com/Azure/go-autorest/autorest/mocks:go_default_library"],
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
        "//vendor/github.com/Azure/go-autorest/autorest/adal/cmd:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
