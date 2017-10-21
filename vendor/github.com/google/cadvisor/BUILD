load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cadvisor.go",
        "storagedriver.go",
    ],
    importpath = "github.com/google/cadvisor",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/cadvisor/cache/memory:go_default_library",
        "//vendor/github.com/google/cadvisor/container:go_default_library",
        "//vendor/github.com/google/cadvisor/http:go_default_library",
        "//vendor/github.com/google/cadvisor/manager:go_default_library",
        "//vendor/github.com/google/cadvisor/storage:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/bigquery:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/elasticsearch:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/influxdb:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/kafka:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/redis:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/statsd:go_default_library",
        "//vendor/github.com/google/cadvisor/storage/stdout:go_default_library",
        "//vendor/github.com/google/cadvisor/utils/sysfs:go_default_library",
        "//vendor/github.com/google/cadvisor/version:go_default_library",
    ],
)

go_binary(
    name = "cadvisor",
    importpath = "github.com/google/cadvisor",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["cadvisor_test.go"],
    importpath = "github.com/google/cadvisor",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/google/cadvisor/container:go_default_library",
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
        "//vendor/github.com/google/cadvisor/api:all-srcs",
        "//vendor/github.com/google/cadvisor/cache:all-srcs",
        "//vendor/github.com/google/cadvisor/collector:all-srcs",
        "//vendor/github.com/google/cadvisor/container:all-srcs",
        "//vendor/github.com/google/cadvisor/devicemapper:all-srcs",
        "//vendor/github.com/google/cadvisor/events:all-srcs",
        "//vendor/github.com/google/cadvisor/fs:all-srcs",
        "//vendor/github.com/google/cadvisor/healthz:all-srcs",
        "//vendor/github.com/google/cadvisor/http:all-srcs",
        "//vendor/github.com/google/cadvisor/info/v1:all-srcs",
        "//vendor/github.com/google/cadvisor/info/v2:all-srcs",
        "//vendor/github.com/google/cadvisor/machine:all-srcs",
        "//vendor/github.com/google/cadvisor/manager:all-srcs",
        "//vendor/github.com/google/cadvisor/metrics:all-srcs",
        "//vendor/github.com/google/cadvisor/pages:all-srcs",
        "//vendor/github.com/google/cadvisor/storage:all-srcs",
        "//vendor/github.com/google/cadvisor/summary:all-srcs",
        "//vendor/github.com/google/cadvisor/utils:all-srcs",
        "//vendor/github.com/google/cadvisor/validate:all-srcs",
        "//vendor/github.com/google/cadvisor/version:all-srcs",
        "//vendor/github.com/google/cadvisor/zfs:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
