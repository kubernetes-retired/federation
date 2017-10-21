load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "client_unix.go",
        "controller.go",
        "event.go",
        "namespace.go",
        "pool.go",
        "rule.go",
        "server_version.go",
        "template.go",
        "util.go",
        "validation.go",
        "volume.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "client_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/storageos/go-api",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gorilla/websocket:go_default_library",
        "//vendor/github.com/storageos/go-api/types:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Microsoft/go-winio:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = [
        "client_test.go",
        "client_unix_test.go",
        "controller_test.go",
        "namespace_test.go",
        "pool_test.go",
        "rule_test.go",
        "storageos_test.go",
        "template_test.go",
        "validation_test.go",
        "volume_test.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "client_windows_test.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/storageos/go-api",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/storageos/go-api/types:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "//vendor/github.com/Microsoft/go-winio:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
        "//vendor/github.com/storageos/go-api/types:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
