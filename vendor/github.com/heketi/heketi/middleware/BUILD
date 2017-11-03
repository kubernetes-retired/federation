load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["jwt.go"],
    importpath = "github.com/heketi/heketi/middleware",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/auth0/go-jwt-middleware:go_default_library",
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/gorilla/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["jwt_test.go"],
    importpath = "github.com/heketi/heketi/middleware",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/dgrijalva/jwt-go:go_default_library",
        "//vendor/github.com/gorilla/context:go_default_library",
        "//vendor/github.com/heketi/heketi/pkg/utils:go_default_library",
        "//vendor/github.com/heketi/tests:go_default_library",
        "//vendor/github.com/urfave/negroni:go_default_library",
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
