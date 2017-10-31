load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["cha.go"],
    importpath = "golang.org/x/tools/go/callgraph/cha",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/callgraph:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
        "//vendor/golang.org/x/tools/go/types/typeutil:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["cha_test.go"],
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/tools/go/callgraph/cha_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/tools/go/callgraph:go_default_library",
        "//vendor/golang.org/x/tools/go/loader:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
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
