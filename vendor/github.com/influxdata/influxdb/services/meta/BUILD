load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "config.go",
        "data.go",
        "errors.go",
        "file_unix.go",
        "query_authorizer.go",
        "write_authorizer.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "file_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/influxdata/influxdb/services/meta",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/influxdata/influxdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/services/meta/internal:go_default_library",
        "//vendor/golang.org/x/crypto/bcrypt:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "data_test.go",
        "meta_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/meta",
    library = ":go_default_library",
    deps = ["//vendor/golang.org/x/crypto/bcrypt:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "client_test.go",
        "config_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/services/meta_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/BurntSushi/toml:go_default_library",
        "//vendor/github.com/influxdata/influxdb:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
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
        "//vendor/github.com/influxdata/influxdb/services/meta/internal:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
