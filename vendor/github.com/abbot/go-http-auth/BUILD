load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "auth.go",
        "basic.go",
        "digest.go",
        "md5crypt.go",
        "misc.go",
        "users.go",
    ],
    importpath = "github.com/abbot/go-http-auth",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "basic_test.go",
        "digest_test.go",
        "md5crypt_test.go",
        "misc_test.go",
        "users_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "github.com/abbot/go-http-auth",
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
