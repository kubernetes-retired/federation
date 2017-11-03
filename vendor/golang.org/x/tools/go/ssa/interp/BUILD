load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "external.go",
        "interp.go",
        "map.go",
        "ops.go",
        "reflect.go",
        "value.go",
    ] + select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "external_darwin.go",
            "external_unix.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "external_unix.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "golang.org/x/tools/go/ssa/interp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/ssa:go_default_library",
        "//vendor/golang.org/x/tools/go/types/typeutil:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            "interp_test.go",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            "interp_test.go",
        ],
        "//conditions:default": [],
    }),
    data = glob(["testdata/**"]),
    importpath = "golang.org/x/tools/go/ssa/interp_test",
    deps = select({
        "@io_bazel_rules_go//go/platform:darwin_amd64": [
            ":go_default_library",
            "//vendor/golang.org/x/tools/go/loader:go_default_library",
            "//vendor/golang.org/x/tools/go/ssa:go_default_library",
            "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
        ],
        "@io_bazel_rules_go//go/platform:linux_amd64": [
            ":go_default_library",
            "//vendor/golang.org/x/tools/go/loader:go_default_library",
            "//vendor/golang.org/x/tools/go/ssa:go_default_library",
            "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
        ],
        "//conditions:default": [],
    }),
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
