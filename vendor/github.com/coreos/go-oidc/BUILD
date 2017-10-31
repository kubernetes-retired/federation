load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "jose.go",
        "jwks.go",
        "oidc.go",
        "verify.go",
    ],
    importpath = "github.com/coreos/go-oidc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pquerna/cachecontrol:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "jwks_test.go",
        "verify_test.go",
    ],
    importpath = "github.com/coreos/go-oidc",
    library = ":go_default_library",
    deps = ["//vendor/gopkg.in/square/go-jose.v2:go_default_library"],
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
        "//vendor/github.com/coreos/go-oidc/example/idtoken:all-srcs",
        "//vendor/github.com/coreos/go-oidc/example/nonce:all-srcs",
        "//vendor/github.com/coreos/go-oidc/example/userinfo:all-srcs",
        "//vendor/github.com/coreos/go-oidc/http:all-srcs",
        "//vendor/github.com/coreos/go-oidc/jose:all-srcs",
        "//vendor/github.com/coreos/go-oidc/key:all-srcs",
        "//vendor/github.com/coreos/go-oidc/oauth2:all-srcs",
        "//vendor/github.com/coreos/go-oidc/oidc:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
