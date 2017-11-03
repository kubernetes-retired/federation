load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["keystore.go"],
    importpath = "github.com/coreos/rkt/pkg/keystore",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["keystore_test.go"],
    importpath = "github.com/coreos/rkt/pkg/keystore",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/coreos/rkt/pkg/keystore/keystoretest:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp/errors:go_default_library",
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
        "//vendor/github.com/coreos/rkt/pkg/keystore/keystoretest:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
