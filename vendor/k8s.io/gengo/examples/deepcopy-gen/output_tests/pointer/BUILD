load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "zz_generated.go",
    ],
    importpath = "k8s.io/gengo/examples/deepcopy-gen/output_tests/pointer",
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/apimachinery/pkg/conversion:go_default_library"],
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
