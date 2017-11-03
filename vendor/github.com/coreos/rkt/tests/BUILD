load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "rkt_flavor.go",
        "rkt_flavor_default.go",
        "rkt_tests.go",
    ],
    importpath = "github.com/coreos/rkt/tests",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/appc/spec/schema:go_default_library",
        "//vendor/github.com/appc/spec/schema/types:go_default_library",
        "//vendor/github.com/coreos/gexpect:go_default_library",
        "//vendor/github.com/coreos/rkt/api/v1alpha:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/tests/testutils:go_default_library",
        "//vendor/github.com/coreos/rkt/tests/testutils/aci-server:go_default_library",
        "//vendor/github.com/kballard/go-shellquote:go_default_library",
        "//vendor/golang.org/x/crypto/openpgp:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)

go_binary(
    name = "tests",
    importpath = "github.com/coreos/rkt/tests",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "rkt_ace_validator_test.go",
        "rkt_app_sandbox_test.go",
        "rkt_attach_test.go",
        "rkt_dns_test.go",
        "rkt_etc_hosts_test.go",
        "rkt_export_container_test.go",
        "rkt_export_test.go",
        "rkt_mount_test.go",
        "rkt_mtab_test.go",
        "rkt_stop_test.go",
    ],
    importpath = "github.com/coreos/rkt/tests",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/coreos/gexpect:go_default_library",
        "//vendor/github.com/coreos/rkt/common:go_default_library",
        "//vendor/github.com/coreos/rkt/tests/testutils:go_default_library",
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
        "//vendor/github.com/coreos/rkt/tests/cniproxy:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/echo-socket-activated:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/inspect:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/rkt-monitor:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/enter:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/gc:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/stub-stage1/run:all-srcs",
        "//vendor/github.com/coreos/rkt/tests/testutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
