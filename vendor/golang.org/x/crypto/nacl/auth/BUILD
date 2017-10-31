load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["auth.go"],
    importpath = "golang.org/x/crypto/nacl/auth",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["auth_test.go"],
    importpath = "golang.org/x/crypto/nacl/auth",
    library = ":go_default_library",
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "golang.org/x/crypto/nacl/auth_test",
    deps = [":go_default_library"],
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
