load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "commune.go",
        "directive.go",
        "package.go",
        "reporting.go",
        "response.go",
        "stressTest.go",
        "stress_client.go",
        "stress_client_query.go",
        "stress_client_write.go",
        "tracer.go",
        "util.go",
    ],
    importpath = "github.com/influxdata/influxdb/stress/v2/stress_client",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/influxdata/influxdb/client/v2:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "commune_test.go",
        "directive_test.go",
        "package_test.go",
        "reporting_test.go",
        "response_test.go",
        "stressTest_test.go",
        "tracer_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/stress/v2/stress_client",
    library = ":go_default_library",
    deps = ["//vendor/github.com/influxdata/influxdb/client/v2:go_default_library"],
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
