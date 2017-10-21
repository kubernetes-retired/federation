load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "attrs.go",
        "client.go",
        "conn.go",
        "packet.go",
        "release.go",
        "server.go",
        "server_statvfs_stubs.go",
        "sftp.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "attrs_unix.go",
            "server_statvfs_darwin.go",
            "server_statvfs_impl.go",
            "server_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "attrs_unix.go",
            "server_statvfs_impl.go",
            "server_statvfs_linux.go",
            "server_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "attrs_stubs.go",
            "server_stubs.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/pkg/sftp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/kr/fs:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "attrs_test.go",
        "client_integration_test.go",
        "client_test.go",
        "other_test.go",
        "packet_test.go",
        "server_integration_test.go",
        "server_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "client_integration_darwin_test.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "client_integration_linux_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/pkg/sftp",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/kr/fs:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "github.com/pkg/sftp_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
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
