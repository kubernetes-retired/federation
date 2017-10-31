load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "importer18.go",
        "importer19.go",
        "stringer.go",
    ],
    importpath = "golang.org/x/tools/cmd/stringer",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "stringer",
    importpath = "golang.org/x/tools/cmd/stringer",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "endtoend_test.go",
        "golden_test.go",
        "util_test.go",
    ],
    importpath = "golang.org/x/tools/cmd/stringer",
    library = ":go_default_library",
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
        "//vendor/golang.org/x/tools/cmd/stringer/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
