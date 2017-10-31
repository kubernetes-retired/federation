load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "asc.go",
        "common.go",
        "doc.go",
        "dockerfetcher.go",
        "downloader.go",
        "fetcher.go",
        "filefetcher.go",
        "finder.go",
        "httpfetcher.go",
        "httpops.go",
        "io.go",
        "namefetcher.go",
        "resumablesession.go",
        "validator.go",
    ],
    importpath = "github.com/coreos/rkt/rkt/image",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/docker2aci/lib:go_default_library",
        "//vendor/github.com/appc/docker2aci/lib/common:go_default_library",
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/discovery:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/ioprogress:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/common/apps:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/distribution:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/keystore:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/log:go_default_library",
        "//vendor/github.com/coreos/rkt/rkt/config:go_default_library",
        "//vendor/github.com/coreos/rkt/rkt/flag:go_default_library",
        "//vendor/github.com/coreos/rkt/rkt/pubkey:go_default_library",
        "//vendor/github.com/coreos/rkt/stage0:go_default_library",
        "//vendor/github.com/coreos/rkt/store/imagestore:go_default_library",
        "//vendor/github.com/coreos/rkt/store/treestore:go_default_library",
        "//vendor/github.com/coreos/rkt/version:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp/errors:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "common_test.go",
        "resumablesession_test.go",
    ],
    importpath = "github.com/coreos/rkt/rkt/image",
    library = ":go_default_library",
    deps = ["//vendor/github.com/coreos/rkt/pkg/distribution:go_default_library"],
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
