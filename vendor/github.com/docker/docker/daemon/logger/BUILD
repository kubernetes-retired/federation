load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "adapter.go",
        "copier.go",
        "factory.go",
        "logger.go",
        "loginfo.go",
        "plugin.go",
        "plugin_unsupported.go",
        "proxy.go",
        "ring.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "plugin_unix.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/docker/docker/daemon/logger",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/backend:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/api/types/plugins/logdriver:go_default_library",
        "//vendor/github.com/docker/docker/pkg/jsonlog:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/github.com/sirupsen/logrus:go_default_library",
    ] + select({
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "//vendor/github.com/tonistiigi/fifo:go_default_library",
            "//vendor/golang.org/x/sys/unix:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

go_test(
    name = "go_default_test",
    srcs = [
        "adapter_test.go",
        "copier_test.go",
        "logger_test.go",
        "ring_test.go",
    ],
    importpath = "github.com/docker/docker/daemon/logger",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/api/types/backend:go_default_library",
        "//vendor/github.com/docker/docker/api/types/plugins/logdriver:go_default_library",
        "//vendor/github.com/gogo/protobuf/io:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
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
        "//vendor/github.com/docker/docker/daemon/logger/awslogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/etwlogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/fluentd:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/gcplogs:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/gelf:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/journald:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/jsonfilelog:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/logentries:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/loggerutils:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/splunk:all-srcs",
        "//vendor/github.com/docker/docker/daemon/logger/syslog:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
