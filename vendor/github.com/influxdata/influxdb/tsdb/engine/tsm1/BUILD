load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bit_reader.go",
        "bool.go",
        "cache.go",
        "compact.go",
        "cursor.go",
        "encoding.gen.go",
        "encoding.go",
        "engine.go",
        "file_store.gen.go",
        "file_store.go",
        "file_unix.go",
        "float.go",
        "int.go",
        "iterator.gen.go",
        "iterator.go",
        "mmap_unix.go",
        "pools.go",
        "reader.go",
        "string.go",
        "timestamp.go",
        "tombstone.go",
        "wal.go",
        "writer.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows_amd64": [
            "file_windows.go",
            "mmap_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "github.com/influxdata/influxdb/tsdb/engine/tsm1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/dgryski/go-bits:go_default_library",
        "//vendor/github.com/dgryski/go-bitstream:go_default_library",
        "//vendor/github.com/golang/snappy:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/pkg/limiter:go_default_library",
        "//vendor/github.com/influxdata/influxdb/pkg/pool:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
        "//vendor/github.com/jwilder/encoding/simple8b:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cache_test.go",
        "int_test.go",
        "string_test.go",
        "timestamp_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/tsdb/engine/tsm1",
    library = ":go_default_library",
    deps = ["//vendor/github.com/golang/snappy:go_default_library"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "bit_reader_test.go",
        "bool_test.go",
        "cache_race_test.go",
        "compact_test.go",
        "encoding_test.go",
        "engine_test.go",
        "file_store_test.go",
        "float_test.go",
        "reader_test.go",
        "tombstone_test.go",
        "wal_test.go",
        "writer_test.go",
    ],
    importpath = "github.com/influxdata/influxdb/tsdb/engine/tsm1_test",
    deps = [
        ":go_default_library",
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/dgryski/go-bitstream:go_default_library",
        "//vendor/github.com/golang/snappy:go_default_library",
        "//vendor/github.com/influxdata/influxdb/influxql:go_default_library",
        "//vendor/github.com/influxdata/influxdb/models:go_default_library",
        "//vendor/github.com/influxdata/influxdb/pkg/deep:go_default_library",
        "//vendor/github.com/influxdata/influxdb/tsdb:go_default_library",
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
