load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cbc_hmac.go",
        "concat_kdf.go",
        "ecdh_es.go",
        "key_wrap.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/cipher",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cbc_hmac_test.go",
        "concat_kdf_test.go",
        "ecdh_es_test.go",
        "key_wrap_test.go",
    ],
    embed = [":go_default_library"],
    importpath = "gopkg.in/square/go-jose.v2/cipher",
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
