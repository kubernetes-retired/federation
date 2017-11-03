load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "a.go",
        "b.go",
        "doc.go",
        "interfaces.go",
        "zz_generated.go",
    ],
    importpath = "k8s.io/gengo/examples/deepcopy-gen/output_tests/wholepkg",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/conversion:go_default_library",
        "//vendor/k8s.io/gengo/examples/deepcopy-gen/output_tests/otherpkg:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["deepcopy_test.go"],
    importpath = "k8s.io/gengo/examples/deepcopy-gen/output_tests/wholepkg",
    library = ":go_default_library",
    deps = ["//vendor/github.com/google/gofuzz:go_default_library"],
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
