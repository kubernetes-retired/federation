load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "bodystring.go",
        "jsonutils.go",
        "log.go",
        "sortedstrings.go",
        "statusgroup.go",
        "stringset.go",
        "stringstack.go",
        "uuid.go",
    ],
    importpath = "github.com/heketi/heketi/pkg/utils",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/lpabon/godbc:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "log_test.go",
        "sortedstrings_test.go",
        "statusgroup_test.go",
        "stringset_test.go",
    ],
    importpath = "github.com/heketi/heketi/pkg/utils",
    library = ":go_default_library",
    deps = ["//vendor/github.com/heketi/tests:go_default_library"],
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
        "//vendor/github.com/heketi/heketi/pkg/utils/ssh:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
