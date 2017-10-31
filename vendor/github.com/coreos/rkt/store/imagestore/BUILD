load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "aciinfo.go",
        "migrate.go",
        "remote.go",
        "schema.go",
        "store.go",
        "utils.go",
    ],
    importpath = "github.com/coreos/rkt/store/imagestore",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/appc/spec/aci:go_default_library",
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/backup:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/lock:go_default_library",
        "//vendor/github.com/coreos/rkt/store/db:go_default_library",
        "//vendor/github.com/hashicorp/errwrap:go_default_library",
        "//vendor/github.com/jonboulle/clockwork:go_default_library",
        "//vendor/github.com/peterbourgon/diskv:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "aciinfo_test.go",
        "main_test.go",
        "migrate_test.go",
        "remote_test.go",
        "store_test.go",
    ],
    importpath = "github.com/coreos/rkt/store/imagestore",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/aci:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/aci/acitest:go_default_library",
        "//vendor/github.com/coreos/rkt/pkg/multicall:go_default_library",
        "//vendor/github.com/coreos/rkt/store/db:go_default_library",
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/jonboulle/clockwork:go_default_library",
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
