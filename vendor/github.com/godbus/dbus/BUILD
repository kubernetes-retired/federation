load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "auth.go",
        "auth_external.go",
        "auth_sha1.go",
        "call.go",
        "conn.go",
        "dbus.go",
        "decoder.go",
        "doc.go",
        "encoder.go",
        "export.go",
        "homedir.go",
        "homedir_dynamic.go",
        "message.go",
        "object.go",
        "sig.go",
        "transport_generic.go",
        "variant.go",
        "variant_lexer.go",
        "variant_parser.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:android": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "conn_darwin.go",
            "transport_darwin.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "conn_other.go",
            "transport_unix.go",
            "transport_unixcred_dragonfly.go",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "conn_other.go",
            "transport_unix.go",
            "transport_unixcred_linux.go",
        ],
        "@io_bazel_rules_go//go/platform:nacl": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "conn_other.go",
            "transport_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "conn_other.go",
        ],
        "//conditions:default": [],
    }),
    cgo = True,
    importpath = "github.com/godbus/dbus",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "conn_test.go",
        "encoder_test.go",
        "examples_test.go",
        "export_test.go",
        "proto_test.go",
        "sig_test.go",
        "transport_unix_test.go",
        "variant_test.go",
    ],
    embed = [":go_default_library"],
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
        "//vendor/github.com/godbus/dbus/introspect:all-srcs",
        "//vendor/github.com/godbus/dbus/prop:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
